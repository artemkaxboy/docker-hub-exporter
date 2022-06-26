/*
 * Copyright (c) 2022 Artem Kolin (https://github.com/artemkaxboy)
 */

package cmd

import (
	exporter "docker-hub-exporter/exporter"
	log "github.com/go-pkgz/lgr"
	"github.com/jessevdk/go-flags"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
	"time"
	"unicode"
)

// ServerCommand set of flags and command for server to start
type ServerCommand struct {
	MetricsPath string `long:"telemetry-path" required:"false" default:"/metrics" description:"Metrics path"`
	Namespaces  string `long:"organisations" env:"ORGS" required:"false" description:"Namespaces to expose metrics for"`
	Images      string `long:"images" env:"IMAGES" required:"false" description:"(list) Images to expose metrics for" env-delim:","`
	Timeout     int    `long:"connection-timeout" env:"CONNECTION_TIMEOUT" required:"false" default:"5" description:"Docker Hub connection timeout in seconds"`

	Port    string `long:"listen-address" env:"BIND_PORT" required:"false" default:":9170" description:"Listen address and port"`
	Retries int    `long:"connection-retries" env:"CONNECTION_TIMEOUT" required:"false" default:"3" description:"Connection retries until failure is raised"`
}

type LegacyOptions struct {
	Organisations string `long:"organisations" required:"false"`
	Images        string `long:"images" required:"false"`
}

// Execute starts server with ServerCommand parameters, entry point for "server" command
func (sc *ServerCommand) Execute(unparsedArgs []string) error {

	log.Printf("[INFO] start `server`")
	log.Printf("[DEBUG] options: %+v", sc)

	newNamespaces := noEmptySplit(removeSpaces(sc.Namespaces), ',')
	newImages := noEmptySplit(removeSpaces(sc.Images), ',')

	dashedArgs := make([]string, len(unparsedArgs))
	for i, attr := range unparsedArgs {
		if strings.HasPrefix(attr, "-") {
			dashedArgs[i] = "-" + attr
		} else {
			dashedArgs[i] = attr
		}
	}

	var legacy LegacyOptions
	p := flags.NewParser(&legacy, flags.None)
	_, err := p.ParseArgs(append([]string{""}, dashedArgs...))
	if err != nil {
		return err
	}

	namespaces := append(newNamespaces, noEmptySplit(removeSpaces(legacy.Organisations), ',')...)
	images := append(newImages, noEmptySplit(removeSpaces(legacy.Images), ',')...)

	if len(namespaces) == 0 && len(images) == 0 {
		log.Printf("[WARN] no namespaces or images specified, use `--organisations` or `--images` flags," +
			" or set `ORGS` or `IMAGES` env variables")
		return nil
	}

	e := exporter.New(
		namespaces,
		images,
		sc.Retries,
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

	err = http.ListenAndServe(sc.Port, nil)
	return err
}

func removeSpaces(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func noEmptySplit(str string, sep rune) []string {
	f := func(c rune) bool {
		return c == sep
	}
	return strings.FieldsFunc(str, f)
}
