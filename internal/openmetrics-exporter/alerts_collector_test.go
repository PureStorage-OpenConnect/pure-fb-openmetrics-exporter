package collectors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"
	"regexp"
	"strings"
	"testing"
)

func TestAlertsCollector(t *testing.T) {

	ropen, _ := os.ReadFile("../../test/data/alerts_open.json")
	rall, _ := os.ReadFile("../../test/data/alerts.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var aopen client.AlertsList
	var aall client.AlertsList
	json.Unmarshal(ropen, &aopen)
	json.Unmarshal(rall, &aall)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlall := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/alerts$`)
		urlopen := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/alerts\?filter=state%3D%27open%27$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
		} else if urlopen.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(ropen))
		} else if urlall.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(rall))
		}
	}))
	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	al := make(map[string]float64)
	for _, a := range aopen.Items {
		al[fmt.Sprintf("%s\n%d\n%s\n%s\n%d\n%s\n%s\n%s",
			strings.Replace(a.Action, "\n", " ", -1),
			a.Code,
			a.ComponentName,
			a.ComponentType,
			a.Created,
			a.KBurl,
			a.Severity,
			strings.Replace(a.Summary, "\n", " ", -1))] += 1
	}
	want := make(map[string]bool)
	for a, n := range al {
		alert := strings.Split(a, "\n")
		want[fmt.Sprintf("label:{name:\"action\" value:\"%s\"} label:{name:\"code\" value:\"%s\"} label:{name:\"component_name\" value:\"%s\"} label:{name:\"component_type\" value:\"%s\"} label:{name:\"created\" value:\"%s\"} label:{name:\"kburl\" value:\"%s\"} label:{name:\"severity\" value:\"%s\"} label:{name:\"summary\" value:\"%s\"} gauge:{value:%g}", alert[0], alert[1], alert[2], alert[3], alert[4], alert[5], alert[6], alert[7], n)] = true
	}
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false, false)
	ac := NewAlertsCollector(c)
	metricsCheck(t, ac, want)
	server.Close()
}
