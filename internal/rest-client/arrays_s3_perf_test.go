package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArraysS3Performance(t *testing.T) {

	s3f, _ := os.ReadFile("../../test/data/arrays_s3performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var as3 ArraysS3PerformanceList
	json.Unmarshal(s3f, &as3)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urls3 := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/s3-specific-performance$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
		} else if urls3.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(s3f))
		}
	}))
	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	t.Run("array_s3_specific_1", func(t *testing.T) {
		defer server.Close()
		c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false, false)
		pl := c.GetArraysS3Performance()
		if diff := cmp.Diff(pl.Items, as3.Items); diff != "" {
			t.Errorf("Mismatch (-want +got):\n%s", diff)
		}
	})
}
