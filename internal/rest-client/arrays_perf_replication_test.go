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

func TestArraysPerformanceReplicationLegacyContinuos(t *testing.T) {
	legacy := []byte(`{
  "continuation_token": null,
  "total_item_count": 1,
  "items": [
    {
      "id": "legacy-id",
      "continuos": {
        "received_bytes_per_sec": 1,
        "transmitted_bytes_per_sec": 2,
        "object_backlog": {
          "put_ops_count": 3,
          "delete_ops_count": 4,
          "other_ops_count": 5,
          "bytes_count": 6
        }
      },
      "aggregate": {
        "received_bytes_per_sec": 7,
        "transmitted_bytes_per_sec": 8
      },
      "periodic": {
        "received_bytes_per_sec": 9,
        "transmitted_bytes_per_sec": 10
      },
      "time": 123
    }
  ]
}`)

	var out ArraysPerformanceReplicationList
	if err := json.Unmarshal(legacy, &out); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if len(out.Items) != 1 {
		t.Fatalf("unexpected items length: %d", len(out.Items))
	}
	cont := out.Items[0].GetContinuous()
	if cont == nil || cont.ObjectBacklog == nil {
		t.Fatalf("expected legacy 'continuos' payload to populate GetContinuous().ObjectBacklog")
	}
	if cont.ObjectBacklog.BytesCount != 6 ||
		cont.ObjectBacklog.DeleteOpsCount != 4 ||
		cont.ObjectBacklog.OtherOpsCount != 5 ||
		cont.ObjectBacklog.PutOpsCount != 3 {
		t.Fatalf("unexpected legacy backlog values: %+v", cont.ObjectBacklog)
	}
}
