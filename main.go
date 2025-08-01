package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	solarPanelGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "solar_panel_output_watts",
			Help: "Solar panel power output (random example)",
		},
		[]string{"panel_id"},
	)

	metricsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "metrics_endpoint_requests_total",
			Help: "Total number of requests to the metrics endpoint",
		},
	)
)
