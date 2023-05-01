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

func TestClientsPerformance(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/clients_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
        var cpl ClientsPerformanceList
        json.Unmarshal(res, &cpl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/clients/performance$`)
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
        t.Run("clients_performance_1", func(t *testing.T) {
            defer server.Close()
	    c := NewRestClient(e, "fake-api-token", "latest", false)
            cp := c.GetClientsPerformance()
	    if diff := cmp.Diff(cp.Items, cpl.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
            }
        })
}
