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

func TestBucketReplicaLinks(t *testing.T) {
	res, _ := os.ReadFile("../../test/data/bucket_replica_links.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")

	var want BucketReplicaLinksList
	_ = json.Unmarshal(res, &want)

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/bucket-replica-links$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
			return
		}
		if valid.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
			return
		}
	}))

	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false, false)
	got := c.GetBucketReplicaLinks()
	if diff := cmp.Diff(got.Items, want.Items); diff != "" {
		t.Errorf("Mismatch items (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(got.Total, want.Total); diff != "" {
		t.Errorf("Mismatch total (-want +got):\n%s", diff)
	}
	server.Close()
}


