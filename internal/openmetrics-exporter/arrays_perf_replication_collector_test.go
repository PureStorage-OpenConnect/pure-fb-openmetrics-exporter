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

func TestArraysPerformanceReplicationCollector(t *testing.T) {
	res, _ := os.ReadFile("../../test/data/arrays_performance_replication.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")

	var arrpr client.ArraysPerformanceReplicationList
	_ = json.Unmarshal(res, &arrpr)

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/performance/replication$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
			return
		}
		if valid.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
			return
		}
	}))

	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false, false)
	prc := NewPerfReplicationCollector(c)

	first := arrpr.Items[0].GetContinuous()
	if first == nil {
		t.Fatalf("expected a continuous replication object in fixture")
	}

	var (
		sumBytes  float64
		sumDelete float64
		sumOther  float64
		sumPut    float64
	)
	for i := range arrpr.Items {
		cont := arrpr.Items[i].GetContinuous()
		if cont == nil || cont.ObjectBacklog == nil {
			continue
		}
		sumBytes += cont.ObjectBacklog.BytesCount
		sumDelete += cont.ObjectBacklog.DeleteOpsCount
		sumOther += cont.ObjectBacklog.OtherOpsCount
		sumPut += cont.ObjectBacklog.PutOpsCount
	}

	want := map[string]bool{
		fmt.Sprintf(
			"%s label:{name:\"dimension\" value:\"transmitted_bytes_per_sec\"} gauge:{value:%g}",
			prc.ThroughputDesc.String(),
			first.TransmittedBytesPerSec,
		): true,
		fmt.Sprintf(
			"%s label:{name:\"dimension\" value:\"received_bytes_per_sec\"} gauge:{value:%g}",
			prc.ThroughputDesc.String(),
			first.ReceivedBytesPerSec,
		): true,
		fmt.Sprintf(
			"%s gauge:{value:%g}",
			prc.ObjectBacklogBytes.String(),
			sumBytes,
		): true,
		fmt.Sprintf(
			"%s gauge:{value:%g}",
			prc.ObjectBacklogDeleteOp.String(),
			sumDelete,
		): true,
		fmt.Sprintf(
			"%s gauge:{value:%g}",
			prc.ObjectBacklogOtherOp.String(),
			sumOther,
		): true,
		fmt.Sprintf(
			"%s gauge:{value:%g}",
			prc.ObjectBacklogPutOp.String(),
			sumPut,
		): true,
	}

	metricsCheckWithDesc(t, prc, want)
	server.Close()
}
