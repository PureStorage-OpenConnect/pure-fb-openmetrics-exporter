package collectors


import (
	"fmt"
	"testing"
        "regexp"
        "strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"os"

	"purestorage/fb-openmetrics-exporter/internal/rest-client"
)

func TestArraysHttpPerfCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/arrays_http_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var arrs client.ArraysHttpPerformanceList
	json.Unmarshal(res, &arrs)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/http-specific-performance$`)
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
	want := make(map[string]bool)
        c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
        for _, p := range arrs.Items {
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"others_per_sec\"} gauge:{value:%g}", p.OthersPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"read_dirs_per_sec\"} gauge:{value:%g}", p.ReadDirsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"read_files_per_sec\"} gauge:{value:%g}", p.ReadFilesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"write_files_per_sec\"} gauge:{value:%g}", p.WriteFilesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_other_op\"} gauge:{value:%g}", p.UsecPerOtherOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_read_dir_op\"} gauge:{value:%g}", p.UsecPerReadDirOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_write_dir_op\"} gauge:{value:%g}", p.UsecPerWriteDirOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_write_file_op\"} gauge:{value:%g}", p.UsecPerWriteFileOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_read_file_op\"} gauge:{value:%g}", p.UsecPerReadFileOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"write_dirs_per_sec\"} gauge:{value:%g}", p.WriteDirsPerSec)] = true
        }
	ac := NewHttpPerfCollector(c)
        metricsCheck(t, ac, want)
        server.Close()
}
