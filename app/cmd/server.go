/*
 * Copyright (c) 2022 Artem Kolin (https://github.com/artemkaxboy)
 */

package cmd

import (
	exporter "docker-hub-exporter/exporter"
	log "github.com/go-pkgz/lgr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

const httpPort = "9170"

// ServerCommand set of flags and command for server to start
type ServerCommand struct {
	MetricsPath string   `long:"metrics-path" env:"METRICS_PATH" required:"false" default:"/metrics" description:"Metrics path"`
	Namespaces  []string `long:"namespace" env:"NAMESPACES" required:"false" description:"(list) Namespaces to expose metrics for" env-delim:","`
	Images      []string `long:"image" env:"IMAGES" required:"false" description:"(list) Images to expose metrics for" env-delim:","`
	Timeout     int      `long:"timeout" env:"TIMEOUT" required:"false" default:"5" description:"Docker Hub connection timeout in seconds"`

	//Retries     int      `long:"retries" env:"RETRIES" required:"false" default:"3" description:"Retries until failure is raised."`
}

// Execute starts server with ServerCommand parameters, entry point for "server" command
func (sc *ServerCommand) Execute(_ []string) error {

	log.Printf("[INFO] start `server`")
	log.Printf("[DEBUG] options: %+v", sc)

	if len(sc.Namespaces) == 0 && len(sc.Images) == 0 {
		log.Printf("[WARN] no namespaces or images specified, use `--namespace` or `--image` flags, or set " +
			"`NAMESPACES` or `IMAGES` env variables")
		return nil
	}

	e := exporter.New(
		sc.Namespaces,
		sc.Images,
		sc.Timeout,
		exporter.WithLogger(log.ToStdLogger(log.Default(), "DEBUG")),
		exporter.WithTimeout(time.Second*time.Duration(sc.Timeout)),
	)

	// Register Metrics from each of the endpoints
	// This invokes the Collect method through the prometheus client libraries.
	prometheus.MustRegister(*e)

	// Setup HTTP handler
	http.Handle(sc.MetricsPath, promhttp.Handler())
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("OK"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`<html>
	<head><title>Docker Hub Exporter</title></head>
	<body>
   		<h1>Prometheus exporter for the Docker Hub</h1>
   		<p>For more information, visit <a href='https://github.com/artemkaxboy/docker-hub-exporter'>GitHub</a></p>
   		<p><a href='` + sc.MetricsPath + `'>Metrics</a></p>
	</body>
</html>`))
	})

	err := http.ListenAndServe(":"+httpPort, nil)
	return err
}
