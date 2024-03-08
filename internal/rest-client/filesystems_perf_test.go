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

func TestFileSystemsPerformance(t *testing.T) {

	ff, _ := os.ReadFile("../../test/data/filesystems_for_perf.json")
	fpf, _ := os.ReadFile("../../test/data/filesystems_perf.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
        var fpl FileSystemsPerformanceList
        json.Unmarshal(fpf, &fpl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        furi := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/file-systems$`)
	        fpuri := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/file-systems/performance$`)
                if r.URL.Path == "/api/api_version" {
                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(http.StatusOK)
                        w.Write([]byte(vers))
                } else if furi.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(ff))
                } else if fpuri.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fpf))
		}
	   }))
        endp := strings.Split(server.URL, "/")
        e := endp[len(endp)-1]
        t.Run("filesystems_performance_all", func(t *testing.T) {
	    c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    f := c.GetFileSystems()
            fl := c.GetFileSystemsPerformance(f, "all")
	    if diff := cmp.Diff(fl.Items, fpl.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        t.Run("filesystems_performance_nfs", func(t *testing.T) {
	    c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    f := c.GetFileSystems()
            fl := c.GetFileSystemsPerformance(f, "NFS")
	    if diff := cmp.Diff(fl.Items, fpl.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        t.Run("filesystems_performance_smb", func(t *testing.T) {
	    c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    f := c.GetFileSystems()
            fl := c.GetFileSystemsPerformance(f, "SMB")
	    if diff := cmp.Diff(fl.Items, fpl.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        server.Close()
}
