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

func TestArraysNfsPerfCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/arrays_http_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var arrs client.ArraysNfsPerformanceList
	json.Unmarshal(res, &arrs)
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
	want := make(map[string]bool)
        c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
        for _, p := range arrs.Items {
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"fsinfos_per_sec\"} gauge:{value:%g}", p.FsinfosPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"fsstats_per_sec\"} gauge:{value:%g}", p.FsstatsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"mkdirs_per_sec\"} gauge:{value:%g}", p.MkdirsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"renames_per_sec\"} gauge:{value:%g}", p.RenamesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"setattrs_per_sec\"} gauge:{value:%g}", p.SetattrsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_link_op\"} gauge:{value:%g}", p.UsecPerLinkOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_mkdir_op\"} gauge:{value:%g}", p.UsecPerMkdirOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_share_metadata_reads_per_sec\"} gauge:{value:%g}", p.AggregateShareMetadataReadsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"links_per_sec\"} gauge:{value:%g}", p.LinksPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"readdirs_per_sec\"} gauge:{value:%g}", p.ReaddirsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"rmdirs_per_sec\"} gauge:{value:%g}", p.RmdirsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_fsstat_op\"} gauge:{value:%g}", p.UsecPerFsstatOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_readdirplus_op\"} gauge:{value:%g}", p.UsecPerReaddirplusOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_remove_op\"} gauge:{value:%g}", p.UsecPerRemoveOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_rename_op\"} gauge:{value:%g}", p.UsecPerRenameOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_file_metadata_reads_per_sec\"} gauge:{value:%g}", p.AggregateFileMetadataReadsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"lookups_per_sec\"} gauge:{value:%g}", p.LookupsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"removes_per_sec\"} gauge:{value:%g}", p.RemovesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_usec_per_file_metadata_create_op\"} gauge:{value:%g}", p.AggregateUsecPerFileMetadataCreateOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_usec_per_other_op\"} gauge:{value:%g}", p.AggregateUsecPerOtherOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_access_op\"} gauge:{value:%g}", p.UsecPerAccessOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_lookup_op\"} gauge:{value:%g}", p.UsecPerLookupOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_symlink_op\"} gauge:{value:%g}", p.UsecPerSymlinkOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"getattrs_per_sec\"} gauge:{value:%g}", p.GetattrsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"readdirpluses_per_sec\"} gauge:{value:%g}", p.ReaddirplusesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"writes_per_sec\"} gauge:{value:%g}", p.WritesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_write_op\"} gauge:{value:%g}", p.UsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"accesses_per_sec\"} gauge:{value:%g}", p.AccessesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"creates_per_sec\"} gauge:{value:%g}", p.CreatesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"pathconfs_per_sec\"} gauge:{value:%g}", p.PathconfsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_rmdir_op\"} gauge:{value:%g}", p.UsecPerRmdirOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_file_metadata_modifies_per_sec\"} gauge:{value:%g}", p.AggregateFileMetadataModifiesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"symlinks_per_sec\"} gauge:{value:%g}", p.SymlinksPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_fsinfo_op\"} gauge:{value:%g}", p.UsecPerFsinfoOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_getattr_op\"} gauge:{value:%g}", p.UsecPerGetattrOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_readlink_op\"} gauge:{value:%g}", p.UsecPerReadlinkOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"reads_per_sec\"} gauge:{value:%g}", p.ReadsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_create_op\"} gauge:{value:%g}", p.UsecPerCreateOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_pathconf_op\"} gauge:{value:%g}", p.UsecPerPathconfOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_read_op\"} gauge:{value:%g}", p.UsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_usec_per_share_metadata_read_op\"} gauge:{value:%g}", p.AggregateUsecPerShareMetadataReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_other_per_sec\"} gauge:{value:%g}", p.AggregateOtherPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"readlinks_per_sec\"} gauge:{value:%g}", p.ReadlinksPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_usec_per_file_metadata_modify_op\"} gauge:{value:%g}", p.AggregateUsecPerFileMetadataModifyOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_usec_per_file_metadata_read_op\"} gauge:{value:%g}", p.AggregateUsecPerFileMetadataReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_readdir_op\"} gauge:{value:%g}", p.UsecPerReaddirOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_setattr_op\"} gauge:{value:%g}", p.UsecPerSetattrOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"aggregate_file_metadata_creates_per_sec\"} gauge:{value:%g}", p.AggregateFileMetadataCreatesPerSec)] = true
	}
	ac := NewNfsPerfCollector(c)
        metricsCheck(t, ac, want)
        server.Close()
}
