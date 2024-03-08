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

func TestArraysSpace(t *testing.T) {

	resa, _ := os.ReadFile("../../test/data/arrays_space_array.json")
	resfs, _ := os.ReadFile("../../test/data/arrays_space_file_system.json")
	reso, _ := os.ReadFile("../../test/data/arrays_space_object_store.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
        var arrsa ArraysSpaceList
        var arrsfs ArraysSpaceList
        var arrso ArraysSpaceList
        json.Unmarshal(resa, &arrsa)
        json.Unmarshal(resfs, &arrsfs)
        json.Unmarshal(reso, &arrso)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        urla := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/space\?type=array$`)
	        urlfs := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/space\?type=file-system$`)
	        urlo := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/space\?type=object-store$`)
                if r.URL.Path == "/api/api_version" {
                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(http.StatusOK)
                        w.Write([]byte(vers))
                } else if urla.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(resa))
                } else if urlfs.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(resfs))
                } else if urlo.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(reso))
		}
	   }))
        endp := strings.Split(server.URL, "/")
        e := endp[len(endp)-1]
        t.Run("arrays_space_1", func(t *testing.T) {
	    c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
            asl := c.GetArraysSpace("array")
	    if diff := cmp.Diff(asl.Items, arrsa.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
                server.Close()
            }
        })
        t.Run("arrays_space_2", func(t *testing.T) {
	    c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
            asfs := c.GetArraysSpace("file-system")
	    if diff := cmp.Diff(asfs.Items, arrsfs.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
                server.Close()
            }
        })
        t.Run("arrays_space_3", func(t *testing.T) {
	    c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
            aso := c.GetArraysSpace("object-store")
	    if diff := cmp.Diff(aso.Items, arrso.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
                server.Close()
            }
        })
        server.Close()
}
