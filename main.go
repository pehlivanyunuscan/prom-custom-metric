package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/gofiber/fiber/v2"
	"main.go/logging"
	"main.go/patterngen"
)

// Panel tanımı
type Panel struct {
	Location       string `json:"location"`
	Brand          string `json:"brand"`
	MaxPower       float64
	Pattern        []float64
	lastPatternDay int
}

// Panel listesi
var panelList = []Panel{
	{Location: "Ankara", Brand: "ABC", MaxPower: 1250, lastPatternDay: -1},
	{Location: "Istanbul", Brand: "DEF", MaxPower: 1400, lastPatternDay: -1},
	{Location: "Izmir", Brand: "GHI", MaxPower: 1600, lastPatternDay: -1},
	{Location: "Bursa", Brand: "JKL", MaxPower: 1500, lastPatternDay: -1},
	{Location: "Adana", Brand: "MNO", MaxPower: 1350, lastPatternDay: -1},
}

var (
	sensorLabels = []string{
		"aku gerilimi",
		"panel gerilimi",
		"sarj akimi",
		"yuk akimi",
		"sicaklik",
		"soc",
		"yuk gucu",
		"panel gucu",
		"yuk durum",
		"aku tipi",
		"sarj durum",
		"kapi bilgisi",
		"sarj gucu",
		"panel akimi",
	}

	roleLabels = []string{
		"role 1",
		"role 2",
		"role 3",
		"role 4",
		"role 5",
		"role 6",
		"role 7",
	}

	once         sync.Once
	sensorGauges map[string]*metrics.Gauge
	// Pattern için global değişkenler
	startMinute = 0    // 00:00
	endMinute   = 1440 // 24*60 = 1440
	// maxPanelGucu     = 1000.0 // örnek için
	// panelGucuPattern []float64
	timeGauge *metrics.Gauge // time_active{time="11_14"} gibi bir etiketle
)

// updatePanelPatternIfNeeded güncel panel gücü desenini günceller
// Eğer gün değiştiyse yeni bir desen oluşturur.
func updatePanelPatternIfNeeded() {
	now := time.Now()
	day := now.YearDay()
	for i := range panelList {
		if panelList[i].lastPatternDay != day {
			// Her panel için farklı seed
			seed := int64(day) + int64(i)*1000
			panelList[i].Pattern = patterngen.GenerateDailyPattern(startMinute, endMinute, panelList[i].MaxPower, seed)
			panelList[i].lastPatternDay = day
		}
	}
}

func initGauges() {
	sensorGauges = make(map[string]*metrics.Gauge)
	// Normal sensörler
	for _, sensor := range sensorLabels {
		if sensor == "panel gucu" {
			// Her panel için ayrı metric
			for _, panel := range panelList {
				key := fmt.Sprintf(`mppt_values{sensor="%s",location="%s",brand="%s"}`, sensor, panel.Location, panel.Brand)
				g := metrics.GetOrCreateGauge(key, nil)
				sensorGauges[key] = g
			}
		} else {
			key := fmt.Sprintf(`mppt_values{sensor="%s"}`, sensor)
			g := metrics.GetOrCreateGauge(key, nil)
			sensorGauges[key] = g
		}
	}
	// Role durumları, sensör olarak "role durumlari" ve ek olarak "role" etiketi
	for _, role := range roleLabels {
		key := fmt.Sprintf(`mppt_values{sensor="role durumlari",role="%s"}`, role)
		g := metrics.GetOrCreateGauge(key, nil)
		sensorGauges[key] = g
	}

	timeGauge = metrics.GetOrCreateGauge(`time_active{time="11_15"}`, nil) // 11 le 15 arasındaki zaman dilimi için gauge oluşturulur
}

func randomValue(sensor string) float64 {
	switch sensor {
	case "aku gerilimi", "panel gerilimi":
		return float64(rand.IntN(1500) + 1200)
	case "sarj akimi", "yuk akimi", "panel akimi":
		return float64(rand.IntN(3000) + 500)
	case "sicaklik":
		return float64(rand.IntN(100))
	case "soc":
		return float64(rand.IntN(101))
	case "yuk gucu", "panel gucu", "sarj gucu":
		return float64(rand.IntN(4000) + 1000)
	case "yuk durum", "aku tipi", "sarj durum":
		return float64(rand.IntN(2)) // 0 veya 1
	case "kapi bilgisi":
		return float64(rand.IntN(3)) // 0, 1 veya 2
	case "role durumlari":
		return float64(rand.IntN(2)) // 0 veya 1
	default:
		return 0.0
	}
}

func main() {
	// Logging seviyesini başlat
	logging.SetLogLevel()

	app := fiber.New()
	once.Do(initGauges)

	app.Get("/metrics", func(c *fiber.Ctx) error {
		updatePanelPatternIfNeeded() // Panel gücü desenini güncelle
		// Sensorlar için değer güncelle
		for _, sensor := range sensorLabels {
			var val float64
			if sensor == "panel gucu" {
				for _, panel := range panelList {
					val = patterngen.GetPatternValueForNow(panel.Pattern, startMinute, endMinute)
					// Her panel için ayrı metric
					key := fmt.Sprintf(`mppt_values{sensor="%s",location="%s",brand="%s"}`, sensor, panel.Location, panel.Brand)
					if g, ok := sensorGauges[key]; ok {
						g.Set(val)
					}
				}
			} else {
				val = randomValue(sensor)
			}
			key := fmt.Sprintf(`mppt_values{sensor="%s"}`, sensor)
			if g, ok := sensorGauges[key]; ok {
				g.Set(val)
			}
		}
		// Role’ler için değer güncelle
		for _, role := range roleLabels {
			val := randomValue("role durumlari")
			key := fmt.Sprintf(`mppt_values{sensor="role durumlari",role="%s"}`, role)
			if g, ok := sensorGauges[key]; ok {
				g.Set(val)
			}
		}

		setTimeMetric() // Saat aralığına göre time_active değerini ayarla

		// logging.LogApp(logging.INFO, "/metrics endpointine istek geldi. IP: %s", c.IP())
		/* logging.LogAudit(
			"anonymous",
			"/metrics",
			c.Method(),
			http.StatusOK,
			c.IP(),
			nil,
			"Panel metrikleri listelendi",
		) */

		c.Set("Content-Type", "text/plain; version=0.0.4")
		metrics.WritePrometheus(c.Context(), true)
		return nil
	})

	logging.LogApp(logging.INFO, "Uygulama başlatıldı")
	if err := app.Listen(":8080"); err != nil {
		logging.LogApp(logging.ERROR, "Sunucu başlatılamadı: %v", err)
	}
}

// setTimeMetric, günün saatine göre time_active değerini ayarlar
// 11:00 - 15:00 saatleri arasında aktif, diğer zamanlarda pasif olarak ayarlanır.
// Bu fonksiyon, her istek geldiğinde çağrılır.
func setTimeMetric() {
	now := time.Now()
	hour := now.Hour()

	// Saat ve dakika değerlerini 11-15 aralığına göre ayarla
	if hour >= 8 && hour < 15 {
		timeGauge.Set(1) // Aktif
	} else {
		timeGauge.Set(0) // Pasif
	}
}
