package main

import (
	"math/rand/v2"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"main.go/logging"
)

var (
	// Sadece sensör değerlerini tutan Gauge
	mpptGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "mppt_values",
			Help: "MPPT sensor values",
		},
		[]string{"sensor", "role"},
	)

	// Counter: /metrics endpointine yapılan istek sayısı
	metricsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "metrics_endpoint_requests_total",
			Help: "Total number of requests to the metrics endpoint",
		},
	)
)

var sensorLabels = []string{
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
	"role durumlari",
}

var roleLabels = []string{
	"role 1",
	"role 2",
	"role 3",
	"role 4",
	"role 5",
	"role 6",
	"role 7",
}

func randomValue(sensor string) float64 {
	// Her sensör için rastgele bir değer döndür
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
	case "yuk durum", "aku tipi", "sarj durum", "role durumlari":
		return float64(rand.IntN(2)) // 0 veya 1
	case "kapi bilgisi":
		return float64(rand.IntN(3)) // 0, 1 veya 2
	default:
		return 0.0 // Bilinmeyen sensör
	}
}
func init() {
	// Metricleri register et
	prometheus.MustRegister(mpptGauge)
	prometheus.MustRegister(metricsCounter)
}

func main() {
	app := fiber.New()

	app.Get("/metrics", func(c *fiber.Ctx) error {
		metricsCounter.Inc()

		// Tüm sensörler için (role boş)
		for _, sensor := range sensorLabels {
			value := randomValue(sensor)
			mpptGauge.WithLabelValues(sensor, "").Set(value)
		}

		// Tüm role'ler için (sensor="role durumlari", role=rol ismi)
		for _, role := range roleLabels {
			value := randomValue("role durumlari")
			mpptGauge.WithLabelValues("role durumlari", role).Set(value)
		}

		// App log
		logging.LogApp(logging.INFO, "/metrics endpointine istek geldi. IP: %s", c.IP())

		// Audit log
		logging.LogAudit(
			"anonymous",                   // user (kimlik doğrulama yoksa "anonymous")
			"/metrics",                    // endpoint
			c.Method(),                    // method
			http.StatusOK,                 // status code
			c.IP(),                        // client ip
			nil,                           // params
			"Panel metrikleri listelendi", // message
		)

		return adaptor.HTTPHandler(promhttp.Handler())(c)
	})

	logging.LogApp(logging.INFO, "Uygulama başlatıldı")
	if err := app.Listen(":8080"); err != nil {
		logging.LogApp(logging.ERROR, "Sunucu başlatılamadı: %v", err)
	}
}
