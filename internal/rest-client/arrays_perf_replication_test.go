package client


import (
	"testing"
        "regexp"
        "strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"os"

	"github.com/google/go-cmp/cmp"
)


func TestArraysPerformanceReplication(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/arrays_performance_replication.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var arrpr ArraysPerformanceReplicationList
	json.Unmarshal(res, &arrpr)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/performance/replication$`)
                if r.URL.Path == "/api/api_version" {
                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
                } else if valid.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
	        }
	   }))
        endp := strings.Split(server.URL, "/")
        e := endp[len(endp)-1]
        t.Run("array_performance_replication_1", func(t *testing.T) {
	    defer server.Close()
            c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false, false)
	    aprl := c.GetArraysPerformanceReplication()
	    if diff := cmp.Diff(aprl.Items, arrpr.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        server.Close()
}
