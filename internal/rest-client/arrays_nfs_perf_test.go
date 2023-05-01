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

func TestArraysNfsPerformance(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/arrays_nfs_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
        var arrsp ArraysNfsPerformanceList
        json.Unmarshal(res, &arrsp)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/nfs-specific-performance$`)
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
        t.Run("arrays_http_performance_1", func(t *testing.T) {
            defer server.Close()
	    c := NewRestClient(e, "fake-api-token", "latest", false)
            apl := c.GetArraysNfsPerformance()
	    if diff := cmp.Diff(apl.Items[0], arrsp.Items[0]); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
            }
        })
}
