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

func TestArraysPerformance(t *testing.T) {

	pall, _ := os.ReadFile("../../test/data/arrays_performance_all.json")
	phttp, _ := os.ReadFile("../../test/data/arrays_performance_http.json")
	pnfs, _ := os.ReadFile("../../test/data/arrays_performance_nfs.json")
	psmb, _ := os.ReadFile("../../test/data/arrays_performance_smb.json")
	ps3, _ := os.ReadFile("../../test/data/arrays_performance_s3.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var aall ArraysPerformanceList
	var ahttp ArraysPerformanceList
	var anfs ArraysPerformanceList
	var asmb ArraysPerformanceList
	var as3 ArraysPerformanceList
	json.Unmarshal(pall, &aall)
	json.Unmarshal(phttp, &ahttp)
	json.Unmarshal(pnfs, &anfs)
	json.Unmarshal(psmb, &asmb)
	json.Unmarshal(ps3, &as3)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        urlall := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/performance\?protocol=all$`)
	        urlhttp := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/performance\?protocol=HTTP$`)
	        urlnfs := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/performance\?protocol=NFS$`)
	        urlsmb := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/performance\?protocol=SMB$`)
	        urls3 := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/performance\?protocol=S3$`)
                if r.URL.Path == "/api/api_version" {
                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
                } else if urlall.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(pall))
                } else if urlhttp.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(phttp))
                } else if urlnfs.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(pnfs))
                } else if urlsmb.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(psmb))
                } else if urls3.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(ps3))
		}
	   }))
        endp := strings.Split(server.URL, "/")
        e := endp[len(endp)-1]
        t.Run("array_perf_all", func(t *testing.T) {
            c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    pl := c.GetArraysPerformance("all")
	    if diff := cmp.Diff(pl.Items, aall.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        t.Run("array_perf_http", func(t *testing.T) {
            c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    pl := c.GetArraysPerformance("HTTP")
	    if diff := cmp.Diff(pl.Items, ahttp.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        t.Run("array_perf_nfs", func(t *testing.T) {
            c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    pl := c.GetArraysPerformance("NFS")
	    if diff := cmp.Diff(pl.Items, anfs.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        t.Run("array_perf_smb", func(t *testing.T) {
            c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    pl := c.GetArraysPerformance("SMB")
	    if diff := cmp.Diff(pl.Items, asmb.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        t.Run("array_perf_s3", func(t *testing.T) {
            c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    pl := c.GetArraysPerformance("S3")
	    if diff := cmp.Diff(pl.Items, as3.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        server.Close()
}
