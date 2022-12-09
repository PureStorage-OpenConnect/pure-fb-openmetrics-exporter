package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	collectors "purestorage/fb-openmetrics-exporter/internal/openmetrics-exporter"
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var version string = "1.0.1"
var debug bool = false

func main() {

	host := flag.String("host", "0.0.0.0", "Address of the exporter")
	port := flag.Int("port", 9491, "Port of the exporter")
	d := flag.Bool("debug", false, "Debug")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	debug = *d
	log.Printf("Start Pure Storage FlashBlade exporter v%s on %s", version, addr)

	http.HandleFunc("/", index)
	http.HandleFunc("/metrics/array", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	http.HandleFunc("/metrics/clients", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	http.HandleFunc("/metrics/usage", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	path := strings.Split(r.URL.Path, "/")
	metrics := ""
	if len(path) == 2 {
		metrics = "all"
	} else {
		metrics = path[2]
		switch metrics {
		case "clients":
		case "array":
		case "usage":
		default:
			metrics = "all"
		}
	}
	endpoint := params.Get("endpoint")
	if endpoint == "" {
		http.Error(w, "Endpoint parameter is missing", http.StatusBadRequest)
		return
	}
	apiver := params.Get("api-version")
	if apiver == "" {
		apiver = "latest"
	}
	authHeader := r.Header.Get("Authorization")
	authFields := strings.Fields(authHeader)
	if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
		http.Error(w, "Target authorization token is missing", http.StatusBadRequest)
		return
	}
	apitoken := authFields[1]

	registry := prometheus.NewRegistry()
	fbclient := client.NewRestClient(endpoint, apitoken, apiver, debug)
	if fbclient.Error != nil {
		http.Error(w, "Error connecting to FlashBlade. Check your management endpoint and/or api token are correct.", http.StatusBadRequest)
		return
	}
	collectors.Collector(context.TODO(), metrics, registry, fbclient)

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
	fbclient.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	msg := `<html>
<body>
<h1>Pure Storage FlashBlade OpenMetrics Exporter</h1>
<table>
    <thead>
        <tr>
        <td>Type</td>
        <td>Endpoint</td>
        <td>GET parameters</td>
        <td>Description</td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Full metrics</td>
            <td><a href="/metrics?endpoint=host">/metrics</a></td>
            <td>endpoint</td>
            <td>All array metrics. Expect slow response time.</td>
        </tr>
        <tr>
            <td>Array metrics</td>
            <td><a href="/metrics/array?endpoint=host">/metrics/array</a></td>
            <td>endpoint</td>
            <td>Provides only array related metrics.</td>
        </tr>
        <tr>
            <td>Client metrics</td>
            <td><a href="/metrics/clients?endpoint=host">/metrics/clients</a></td>
            <td>endpoint</td>
            <td>Provides only client related metrics. This is the most time expensive query</td>
        </tr>
        <tr>
            <td>Quota metrics</td>
            <td><a href="/metrics/usage?endpoint=host">/metrics/usage</a></td>
            <td>endpoint</td>
            <td>Provides only quota related metrics.</td>
        </tr>
    </tbody>
</table>
</body>
</html>`

	fmt.Fprintf(w, "%s", msg)
}
