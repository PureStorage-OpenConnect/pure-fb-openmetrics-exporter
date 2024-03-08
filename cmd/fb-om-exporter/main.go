package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	config "purestorage/fb-openmetrics-exporter/internal/config"
	collectors "purestorage/fb-openmetrics-exporter/internal/openmetrics-exporter"
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v3"
)

var version string = "development"
var debug bool = false
var arraytokens config.FlashBladeList

func fileExists(args []string) error {
	_, err := os.Stat(args[0])
	return err
}

func isFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {

	parser := argparse.NewParser("pure-fb-om-exporter", "Pure Storage FB OpenMetrics exporter")
	host := parser.String("a", "address", &argparse.Options{Required: false, Help: "IP address for this exporter to bind to", Default: "0.0.0.0"})
	port := parser.Int("p", "port", &argparse.Options{Required: false, Help: "Port for this exporter to listen", Default: 9491})
	d := parser.Flag("d", "debug", &argparse.Options{Required: false, Help: "Enable debug", Default: false})
	at := parser.File("t", "tokens", os.O_RDONLY, 0600, &argparse.Options{Required: false, Validate: fileExists, Help: "API token(s) map file"})
	cert := parser.String("c", "cert", &argparse.Options{Required: false, Help: "SSL/TLS certificate file. Required only for TLS"})
	key := parser.String("k", "key", &argparse.Options{Required: false, Help: "SSL/TLS private key file. Required only for TLS"})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatalf("Error in token file: %v", err)
	}
	if !isNilFile(*at) {
		defer at.Close()
		buf := make([]byte, 1024)
		arrlist := ""
		for {
			n, err := at.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Reading token file: %v", err)
			}
			if n > 0 {
				arrlist = arrlist + string(buf[:n])
			}
		}
		buf = []byte(arrlist)
		err := yaml.Unmarshal(buf, &arraytokens)
		if err != nil {
			log.Fatalf("Unmarshalling token file: %v", err)
		}
	}
	if (len(*cert) > 0 && len(*key) == 0) || (len(*cert) == 0 && len(*key) > 0) {
		log.Fatal("Both certificate and key must be specified to enable TLS")
	}
	if len(*cert) > 0 && len(*key) > 0 {
		if !isFile(*cert) {
			log.Fatal("TLS cert file not found")
		} else if !isFile(*key) {
			log.Fatal("TLS key file not found")
		}
	}
	debug = *d
	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Start Pure FlashBlade exporter %s on %s", version, addr)

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
	http.HandleFunc("/metrics/policies", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	if isFile(*cert) && isFile(*key) {
		log.Fatal(http.ListenAndServeTLS(addr, *cert, *key, nil))
	} else {
		log.Fatal(http.ListenAndServe(addr, nil))
	}
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
		case "policies":
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
	address, apitoken := arraytokens.GetArrayParams(endpoint)
	if len(authFields) == 2 && strings.ToLower(authFields[0]) == "bearer" {
		apitoken = authFields[1]
		address = endpoint
	}
	if apitoken == "" {
		http.Error(w, "Target authorization token is missing", http.StatusBadRequest)
		return
	}

	uagent := r.Header.Get("User-Agent")
	registry := prometheus.NewRegistry()
	fbclient := client.NewRestClient(address, apitoken, apiver, uagent, debug)
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
	<tr>
            <td>NFS export policies metrics</td>
            <td><a href="/metrics/policies?endpoint=host">/metrics/policies</a></td>
            <td>endpoint</td>
            <td>Provides only NFS policies related metrics.</td>
        </tr>
    </tbody>
</table>
</body>
</html>`

	fmt.Fprintf(w, "%s", msg)
}

func isNilFile(f os.File) bool {
	var tf os.File
	return f == tf
}
