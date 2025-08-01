package main

import (
	"math/rand/v2"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Gauge: güneş paneli çıkış gücü (panel_id ile label)
	solarPanelGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "solar_panel_output_watts",
			Help: "Solar panel power output (random example)",
		},
		[]string{"panel_id"},
	)
	// Counter: /metrics endpointine yapılan istek sayısı
	metricsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "metrics_endpoint_requests_total",
			Help: "Total number of requests to the metrics endpoint",
		},
	)
)

func init() {
	// Metricleri register et
	prometheus.MustRegister(solarPanelGauge)
	prometheus.MustRegister(metricsCounter)
}

func main() {
	app := fiber.New()

	app.Get("/metrics", func(c *fiber.Ctx) error {
		metricsCounter.Inc()

		for _, panelID := range []string{"A", "B", "C", "D", "E", "F"} {
			value := 100 + rand.Float64()*400 // Örnek değer
			solarPanelGauge.WithLabelValues(panelID).Set(value)
		}

		return adaptor.HTTPHandler(promhttp.Handler())(c)
	})

	app.Listen(":8080")
}
