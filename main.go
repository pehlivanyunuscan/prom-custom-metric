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

// Panel tanımı - il, ilçe, mahalle eklendi
type Panel struct {
	City           string  `json:"city"`
	District       string  `json:"district"`
	Neighborhood   string  `json:"neighborhood"`
	Brand          string  `json:"brand"`
	MaxPower       float64 `json:"max_power"`
	Pattern        []float64
	lastPatternDay int
}

// Panel listesi
var panelList = []Panel{
	{City: "Ankara", District: "Cankaya", Neighborhood: "Bahcelievler", Brand: "ABC", MaxPower: 1250, lastPatternDay: -1},
	{City: "Ankara", District: "Yenimahalle", Neighborhood: "Demetevler", Brand: "DEF", MaxPower: 1280, lastPatternDay: -1},
	{City: "Ankara", District: "Mamak", Neighborhood: "Ege", Brand: "GHI", MaxPower: 1320, lastPatternDay: -1},
	{City: "Ankara", District: "Kecioren", Neighborhood: "Etlik", Brand: "JKL", MaxPower: 1270, lastPatternDay: -1},
	{City: "Ankara", District: "Altindag", Neighborhood: "Aydinlikevler", Brand: "MNO", MaxPower: 1300, lastPatternDay: -1},
	{City: "Ankara", District: "Sincan", Neighborhood: "Fatih", Brand: "PQR", MaxPower: 1310, lastPatternDay: -1},
	{City: "Ankara", District: "Golbasi", Neighborhood: "Safak", Brand: "STU", MaxPower: 1240, lastPatternDay: -1},
	{City: "Ankara", District: "Etimesgut", Neighborhood: "Eryaman", Brand: "VWX", MaxPower: 1290, lastPatternDay: -1},

	{City: "Istanbul", District: "Kadikoy", Neighborhood: "Moda", Brand: "DEF", MaxPower: 1400, lastPatternDay: -1},
	{City: "Istanbul", District: "Besiktas", Neighborhood: "Levent", Brand: "ABC", MaxPower: 1450, lastPatternDay: -1},
	{City: "Istanbul", District: "Sisli", Neighborhood: "Nisantasi", Brand: "GHI", MaxPower: 1410, lastPatternDay: -1},
	{City: "Istanbul", District: "Uskudar", Neighborhood: "Cengelkoy", Brand: "JKL", MaxPower: 1430, lastPatternDay: -1},
	{City: "Istanbul", District: "Fatih", Neighborhood: "Aksaray", Brand: "MNO", MaxPower: 1390, lastPatternDay: -1},
	{City: "Istanbul", District: "Bakirkoy", Neighborhood: "Yesilkoy", Brand: "PQR", MaxPower: 1460, lastPatternDay: -1},
	{City: "Istanbul", District: "Beyoglu", Neighborhood: "Cihangir", Brand: "STU", MaxPower: 1420, lastPatternDay: -1},
	{City: "Istanbul", District: "Maltepe", Neighborhood: "Kucukyali", Brand: "VWX", MaxPower: 1440, lastPatternDay: -1},

	{City: "Izmir", District: "Konak", Neighborhood: "Alsancak", Brand: "GHI", MaxPower: 1600, lastPatternDay: -1},
	{City: "Izmir", District: "Bornova", Neighborhood: "Kazimdirik", Brand: "ABC", MaxPower: 1620, lastPatternDay: -1},
	{City: "Izmir", District: "Karsiyaka", Neighborhood: "Bostanli", Brand: "DEF", MaxPower: 1580, lastPatternDay: -1},
	{City: "Izmir", District: "Balcova", Neighborhood: "Teleferik", Brand: "JKL", MaxPower: 1550, lastPatternDay: -1},
	{City: "Izmir", District: "Gaziemir", Neighborhood: "Sarnic", Brand: "PQR", MaxPower: 1570, lastPatternDay: -1},
	{City: "Izmir", District: "Cigli", Neighborhood: "Atasehir", Brand: "MNO", MaxPower: 1610, lastPatternDay: -1},
	{City: "Izmir", District: "Buca", Neighborhood: "Sirinyer", Brand: "STU", MaxPower: 1540, lastPatternDay: -1},
	{City: "Izmir", District: "Narlidere", Neighborhood: "Ilica", Brand: "VWX", MaxPower: 1590, lastPatternDay: -1},

	{City: "Bursa", District: "Osmangazi", Neighborhood: "Heykel", Brand: "JKL", MaxPower: 1500, lastPatternDay: -1},
	{City: "Bursa", District: "Nilufer", Neighborhood: "Gorukle", Brand: "ABC", MaxPower: 1520, lastPatternDay: -1},
	{City: "Bursa", District: "Yildirim", Neighborhood: "Ertugrulgazi", Brand: "DEF", MaxPower: 1510, lastPatternDay: -1},
	{City: "Bursa", District: "Inegol", Neighborhood: "Alanyurt", Brand: "GHI", MaxPower: 1490, lastPatternDay: -1},
	{City: "Bursa", District: "Gemlik", Neighborhood: "Kumla", Brand: "MNO", MaxPower: 1505, lastPatternDay: -1},
	{City: "Bursa", District: "Mudanya", Neighborhood: "Halitpasa", Brand: "PQR", MaxPower: 1515, lastPatternDay: -1},
	{City: "Bursa", District: "Kestel", Neighborhood: "Barakfakih", Brand: "STU", MaxPower: 1530, lastPatternDay: -1},
	{City: "Bursa", District: "Karacabey", Neighborhood: "Yenikoy", Brand: "VWX", MaxPower: 1480, lastPatternDay: -1},

	{City: "Adana", District: "Seyhan", Neighborhood: "Resatbey", Brand: "MNO", MaxPower: 1350, lastPatternDay: -1},
	{City: "Adana", District: "Cukurova", Neighborhood: "Turgut Ozal", Brand: "ABC", MaxPower: 1370, lastPatternDay: -1},
	{City: "Adana", District: "Yuregir", Neighborhood: "Serinevler", Brand: "DEF", MaxPower: 1340, lastPatternDay: -1},
	{City: "Adana", District: "Ceyhan", Neighborhood: "Buyukmangit", Brand: "GHI", MaxPower: 1360, lastPatternDay: -1},
	{City: "Adana", District: "Kozan", Neighborhood: "Sevkiye", Brand: "JKL", MaxPower: 1330, lastPatternDay: -1},
	{City: "Adana", District: "Saricam", Neighborhood: "Carkipare", Brand: "PQR", MaxPower: 1320, lastPatternDay: -1},
	{City: "Adana", District: "Karatas", Neighborhood: "Deniz", Brand: "STU", MaxPower: 1380, lastPatternDay: -1},
	{City: "Adana", District: "Pozanti", Neighborhood: "Kamisli", Brand: "VWX", MaxPower: 1310, lastPatternDay: -1},

	{City: "Antalya", District: "Muratpasa", Neighborhood: "Lara", Brand: "PQR", MaxPower: 1450, lastPatternDay: -1},
	{City: "Antalya", District: "Kepez", Neighborhood: "Varsak", Brand: "ABC", MaxPower: 1430, lastPatternDay: -1},
	{City: "Antalya", District: "Konyaalti", Neighborhood: "Uncali", Brand: "DEF", MaxPower: 1420, lastPatternDay: -1},
	{City: "Antalya", District: "Alanya", Neighborhood: "Mahmutlar", Brand: "GHI", MaxPower: 1460, lastPatternDay: -1},
	{City: "Antalya", District: "Manavgat", Neighborhood: "Side", Brand: "JKL", MaxPower: 1440, lastPatternDay: -1},
	{City: "Antalya", District: "Serik", Neighborhood: "Belek", Brand: "MNO", MaxPower: 1470, lastPatternDay: -1},
	{City: "Antalya", District: "Kumluca", Neighborhood: "Mavikent", Brand: "STU", MaxPower: 1480, lastPatternDay: -1},
	{City: "Antalya", District: "Finike", Neighborhood: "Hasyurt", Brand: "VWX", MaxPower: 1410, lastPatternDay: -1},

	{City: "Konya", District: "Meram", Neighborhood: "Yazir", Brand: "STU", MaxPower: 1550, lastPatternDay: -1},
	{City: "Konya", District: "Selcuklu", Neighborhood: "Bosna Hersek", Brand: "ABC", MaxPower: 1580, lastPatternDay: -1},
	{City: "Konya", District: "Karatay", Neighborhood: "Aziziye", Brand: "DEF", MaxPower: 1540, lastPatternDay: -1},
	{City: "Konya", District: "Beysehir", Neighborhood: "Huyuk", Brand: "GHI", MaxPower: 1520, lastPatternDay: -1},
	{City: "Konya", District: "Seydisehir", Neighborhood: "Kavak", Brand: "JKL", MaxPower: 1560, lastPatternDay: -1},
	{City: "Konya", District: "Cihanbeyli", Neighborhood: "Taspinar", Brand: "MNO", MaxPower: 1530, lastPatternDay: -1},
	{City: "Konya", District: "Cumra", Neighborhood: "Icericumra", Brand: "PQR", MaxPower: 1570, lastPatternDay: -1},
	{City: "Konya", District: "Aksehir", Neighborhood: "Tipi", Brand: "VWX", MaxPower: 1590, lastPatternDay: -1},

	{City: "Trabzon", District: "Ortahisar", Neighborhood: "Merkez", Brand: "VWX", MaxPower: 1300, lastPatternDay: -1},
	{City: "Trabzon", District: "Akcaabat", Neighborhood: "Sogutlu", Brand: "ABC", MaxPower: 1330, lastPatternDay: -1},
	{City: "Trabzon", District: "Yomra", Neighborhood: "Kasustu", Brand: "DEF", MaxPower: 1310, lastPatternDay: -1},
	{City: "Trabzon", District: "Vakfikebir", Neighborhood: "Carsi", Brand: "GHI", MaxPower: 1320, lastPatternDay: -1},
	{City: "Trabzon", District: "Arakli", Neighborhood: "Yaliboyu", Brand: "JKL", MaxPower: 1290, lastPatternDay: -1},
	{City: "Trabzon", District: "Macka", Neighborhood: "Istiklal", Brand: "MNO", MaxPower: 1340, lastPatternDay: -1},
	{City: "Trabzon", District: "Surmene", Neighborhood: "Yeniay", Brand: "PQR", MaxPower: 1350, lastPatternDay: -1},
	{City: "Trabzon", District: "Besikduzu", Neighborhood: "Fatih", Brand: "STU", MaxPower: 1280, lastPatternDay: -1},

	{City: "Mugla", District: "Bodrum", Neighborhood: "Yalikavak", Brand: "XYZ", MaxPower: 1200, lastPatternDay: -1},
	{City: "Mugla", District: "Fethiye", Neighborhood: "Calis", Brand: "ABC", MaxPower: 1100, lastPatternDay: -1},
	{City: "Mugla", District: "Mentese", Neighborhood: "Orhaniye", Brand: "DEF", MaxPower: 1230, lastPatternDay: -1},
	{City: "Mugla", District: "Milas", Neighborhood: "Gulluk", Brand: "GHI", MaxPower: 1250, lastPatternDay: -1},
	{City: "Mugla", District: "Ortaca", Neighborhood: "Dalyan", Brand: "JKL", MaxPower: 1210, lastPatternDay: -1},
	{City: "Mugla", District: "Datca", Neighborhood: "Resadiye", Brand: "MNO", MaxPower: 1190, lastPatternDay: -1},
	{City: "Mugla", District: "Ula", Neighborhood: "Akyaka", Brand: "PQR", MaxPower: 1170, lastPatternDay: -1},
	{City: "Mugla", District: "Koycegiz", Neighborhood: "Toparlar", Brand: "STU", MaxPower: 1240, lastPatternDay: -1},

	{City: "Duzce", District: "Merkez", Neighborhood: "Serefiye", Brand: "DEF", MaxPower: 1150, lastPatternDay: -1},
	{City: "Duzce", District: "Akcakoca", Neighborhood: "Osmaniye", Brand: "ABC", MaxPower: 1160, lastPatternDay: -1},
	{City: "Duzce", District: "Cumayeri", Neighborhood: "Cilimli", Brand: "GHI", MaxPower: 1130, lastPatternDay: -1},
	{City: "Duzce", District: "Golyaka", Neighborhood: "Yazlik", Brand: "JKL", MaxPower: 1140, lastPatternDay: -1},
	{City: "Duzce", District: "Kaynasli", Neighborhood: "Camlica", Brand: "MNO", MaxPower: 1170, lastPatternDay: -1},
	{City: "Duzce", District: "Yigilca", Neighborhood: "Aksu", Brand: "PQR", MaxPower: 1180, lastPatternDay: -1},
	{City: "Duzce", District: "Gumusova", Neighborhood: "Yazlik", Brand: "STU", MaxPower: 1120, lastPatternDay: -1},
	{City: "Duzce", District: "Cilimli", Neighborhood: "Yenikoy", Brand: "VWX", MaxPower: 1190, lastPatternDay: -1},
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
				key := fmt.Sprintf(`mppt_values{sensor="%s",city="%s",district="%s",neighborhood="%s",brand="%s"}`, sensor, panel.City, panel.District, panel.Neighborhood, panel.Brand)
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
					key := fmt.Sprintf(`mppt_values{sensor="%s",city="%s",district="%s",neighborhood="%s",brand="%s"}`, sensor, panel.City, panel.District, panel.Neighborhood, panel.Brand)
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
	if hour >= 8 && hour < 17 {
		timeGauge.Set(1) // Aktif
	} else {
		timeGauge.Set(0) // Pasif
	}
}
