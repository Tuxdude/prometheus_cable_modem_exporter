package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	listenAddrFormat  = "%s:%d"
	metricsURLFormat  = "http://%s:%d/%s"
	metricsPathFormat = "/%s"
)

func promHandler(cmCollector *collector) http.Handler {
	// Use a custom registry to get rid of the default set of metrics
	// added by prometheus and have full control.
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(
		// Process information collector.
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		// Go Build Information.
		collectors.NewBuildInfoCollector(),
		// Cable Modem metrics collector.
		cmCollector,
		// Disable Go process metrics collector.
		// collectors.NewGoCollector(),
	)
	return promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.HTTPErrorOnError,
		},
	)
}

func startExporter(
	listenHost string,
	listenPort uint32,
	path string,
	cmCollector *collector,
) {
	logger := buildLogger()
	defer logger.Sync() // nolint - flushes buffer, if any
	log := logger.Sugar()

	listenAddr := fmt.Sprintf(listenAddrFormat, listenHost, listenPort)
	metricsPath := fmt.Sprintf(metricsPathFormat, path)
	metricsURL := fmt.Sprintf(metricsURLFormat, listenHost, listenPort, path)

	// Set up HTTP handler for metrics.
	mux := http.NewServeMux()
	mux.Handle(metricsPath, promHandler(cmCollector))

	// Start listening for HTTP connections.
	server := http.Server{
		Addr:         listenAddr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Infof("Starting Cable Modem prometheus metrics exporter on %s", metricsURL)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start Cable Modem prometheus metrics exporter: %s", err)
	}
}
