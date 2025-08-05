package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"sync"

	"github.com/VictoriaMetrics/metrics"
	"github.com/gofiber/fiber/v2"
	"main.go/logging"
)

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
	roleGauges   map[string]map[string]*metrics.Gauge
)

func initGauges() {
	sensorGauges = make(map[string]*metrics.Gauge)
	for _, sensor := range sensorLabels {
		g := metrics.GetOrCreateGauge(fmt.Sprintf(`mppt_values{sensor="%s"}`, sensor), nil)
		sensorGauges[sensor] = g
	}

	roleGauges = make(map[string]map[string]*metrics.Gauge)
	roleGauges["role durumlari"] = make(map[string]*metrics.Gauge)
	for _, role := range roleLabels {
		// sensor="role durumlari", role="role x"
		key := role
		g := metrics.GetOrCreateGauge(fmt.Sprintf(`mppt_values{sensor="role durumlari",role="%s"}`, key), nil)
		roleGauges["role durumlari"][key] = g
	}
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
	app := fiber.New()

	once.Do(initGauges)

	app.Get("/metrics", func(c *fiber.Ctx) error {
		// Sensorlar için değer güncelle
		for _, sensor := range sensorLabels {
			if sensor == "role durumlari" {
				continue // Role durumları için ayrı işlem yapacağız
			}
			val := randomValue(sensor)
			if g, ok := sensorGauges[sensor]; ok {
				g.Set(val)
			}
		}
		// Role’ler için değer güncelle
		for _, role := range roleLabels {
			val := randomValue("role durumlari")
			if g, ok := roleGauges["role durumlari"][role]; ok {
				g.Set(val)
			}
		}

		logging.LogApp(logging.INFO, "/metrics endpointine istek geldi. IP: %s", c.IP())
		logging.LogAudit(
			"anonymous",
			"/metrics",
			c.Method(),
			http.StatusOK,
			c.IP(),
			nil,
			"Panel metrikleri listelendi",
		)

		c.Set("Content-Type", "text/plain; version=0.0.4")
		metrics.WritePrometheus(c.Context(), true)
		return nil
	})

	logging.LogApp(logging.INFO, "Uygulama başlatıldı")
	if err := app.Listen(":8080"); err != nil {
		logging.LogApp(logging.ERROR, "Sunucu başlatılamadı: %v", err)
	}
}
