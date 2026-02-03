package collectors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"
	"regexp"
	"strings"
	"testing"

)

func TestBucketReplicaLinksCollector(t *testing.T) {
	res, _ := os.ReadFile("../../test/data/bucket_replica_links.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")

	var parsed client.BucketReplicaLinksList
	_ = json.Unmarshal(res, &parsed)

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
	fb := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false, false)
	c := NewBucketReplicaLinksCollector(fb)

	want := make(map[string]bool)
	for i := range parsed.Items {
		item := parsed.Items[i]
		lb := item.LocalBucket.Name
		remote := item.Remote.Name
		dir := item.Direction

		want[fmt.Sprintf("%s label:{name:\"direction\" value:\"%s\"} label:{name:\"local_bucket\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} gauge:{value:%g}", c.LagDesc.String(), dir, lb, remote, item.Lag)] = true

		want[fmt.Sprintf("%s label:{name:\"direction\" value:\"%s\"} label:{name:\"local_bucket\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} gauge:{value:%g}", c.ObjectBacklogBytesDesc.String(), dir, lb, remote, item.ObjectBacklog.BytesCount)] = true
		want[fmt.Sprintf("%s label:{name:\"direction\" value:\"%s\"} label:{name:\"local_bucket\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} gauge:{value:%g}", c.ObjectBacklogDeleteOpsDesc.String(), dir, lb, remote, item.ObjectBacklog.DeleteOpsCount)] = true
		want[fmt.Sprintf("%s label:{name:\"direction\" value:\"%s\"} label:{name:\"local_bucket\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} gauge:{value:%g}", c.ObjectBacklogOtherOpsDesc.String(), dir, lb, remote, item.ObjectBacklog.OtherOpsCount)] = true
		want[fmt.Sprintf("%s label:{name:\"direction\" value:\"%s\"} label:{name:\"local_bucket\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} gauge:{value:%g}", c.ObjectBacklogPutOpsDesc.String(), dir, lb, remote, item.ObjectBacklog.PutOpsCount)] = true
	}

	if parsed.Total == nil || parsed.Total.ObjectBacklog == nil {
		t.Fatalf("expected fixture to include total.object_backlog")
	}
	want[fmt.Sprintf("%s gauge:{value:%g}", c.TotalLagDesc.String(), parsed.Total.Lag)] = true
	want[fmt.Sprintf("%s gauge:{value:%g}", c.TotalObjectBacklogBytesDesc.String(), parsed.Total.ObjectBacklog.BytesCount)] = true
	want[fmt.Sprintf("%s gauge:{value:%g}", c.TotalObjectBacklogDeleteOpsDesc.String(), parsed.Total.ObjectBacklog.DeleteOpsCount)] = true
	want[fmt.Sprintf("%s gauge:{value:%g}", c.TotalObjectBacklogOtherOpsDesc.String(), parsed.Total.ObjectBacklog.OtherOpsCount)] = true
	want[fmt.Sprintf("%s gauge:{value:%g}", c.TotalObjectBacklogPutOpsDesc.String(), parsed.Total.ObjectBacklog.PutOpsCount)] = true

	metricsCheckWithDesc(t, c, want)
	server.Close()
}


