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

func TestBucketsPerformance(t *testing.T) {

	bf, _ := os.ReadFile("../../test/data/buckets_for_perf.json")
	bpf, _ := os.ReadFile("../../test/data/buckets_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
        var bpl BucketsPerformanceList
        json.Unmarshal(bpf, &bpl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        buri := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/buckets$`)
	        bpuri := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/buckets/performance$`)
                if r.URL.Path == "/api/api_version" {
                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(http.StatusOK)
                        w.Write([]byte(vers))
                } else if buri.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(bf))
                } else if bpuri.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(bpf))
		}
	   }))
        endp := strings.Split(server.URL, "/")
        e := endp[len(endp)-1]
        t.Run("buckets_performance_1", func(t *testing.T) {
            defer server.Close()
	    c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    b := c.GetBuckets()
            pl := c.GetBucketsPerformance(b)
	    if diff := cmp.Diff(pl.Items, bpl.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
            }
        })
}
