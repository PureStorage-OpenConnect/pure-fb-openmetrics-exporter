package restclient

type ArrayNfsPerformance struct {
    Name                                 string   `json:"name"`
    Id                                   string `json:"id"`
    AccessesPerSec                       float64 `json:"accesses_per_sec"`
    AggregateFileMetadataCreatesPerSec   float64 `json:"aggregate_file_metadata_creates_per_sec"`
    AggregateFileMetadataModifiesPerSec  float64 `json:"aggregate_file_metadata_modifies_per_sec"`
    AggregateFileMetadataReadsPerSec     float64 `json:"aggregate_file_metadata_reads_per_sec"`
    AggregateOtherPerSec                 float64 `json:"aggregate_other_per_sec"`
    AggregateShareMetadataReadsPerSec    float64 `json:"aggregate_share_metadata_reads_per_sec"`
    AggregateUsecPerFileMetadataCreateOp float64 `json:"aggregate_usec_per_file_metadata_create_op"`
    AggregateUsecPerFileMetadataModifyOp float64 `json:"aggregate_usec_per_file_metadata_modify_op"`
    AggregateUsecPerFileMetadataReadOp   float64 `json:"aggregate_usec_per_file_metadata_read_op"`
    AggregateUsecPerOtherOp              float64 `json:"aggregate_usec_per_other_op"`
    AggregateUsecPerShareMetadataReadOp  float64 `json:"aggregate_usec_per_share_metadata_read_op"`
    CreatesPerSec                        float64 `json:"creates_per_sec"`
    FsinfosPerSec                        float64 `json:"fsinfos_per_sec"`
    FsstatsPerSec                        float64 `json:"fsstats_per_sec"`
    GetattrsPerSec                       float64 `json:"getattrs_per_sec"`
    LinksPerSec                          float64 `json:"links_per_sec"`
    LookupsPerSec                        float64 `json:"lookups_per_sec"`
    MkdirsPerSec                         float64 `json:"mkdirs_per_sec"`
    PathconfsPerSec                      float64 `json:"pathconfs_per_sec"`
    ReadsPerSec                          float64 `json:"reads_per_sec"`
    ReaddirsPerSec                       float64 `json:"readdirs_per_sec"`
    ReaddirplusesPerSec                  float64 `json:"readdirpluses_per_sec"`
    ReadlinksPerSec                      float64 `json:"readlinks_per_sec"`
    RemovesPerSec                        float64 `json:"removes_per_sec"`
    RenamesPerSec                        float64 `json:"renames_per_sec"`
    RmdirsPerSec                         float64 `json:"rmdirs_per_sec"`
    SetattrsPerSec                       float64 `json:"setattrs_per_sec"`
    SymlinksPerSec                       float64 `json:"symlinks_per_sec"`
    Time                                 int     `json:"time"`
    WritesPerSec                         float64 `json:"writes_per_sec"`
    UsecPerAccessOp                      float64 `json:"usec_per_access_op"`
    UsecPerCreateOp                      float64 `json:"usec_per_create_op"`
    UsecPerFsinfoOp                      float64 `json:"usec_per_fsinfo_op"`
    UsecPerFsstatOp                      float64 `json:"usec_per_fsstat_op"`
    UsecPerGetattrOp                     float64 `json:"usec_per_getattr_op"`
    UsecPerLinkOp                        float64 `json:"usec_per_link_op"`
    UsecPerLookupOp                      float64 `json:"usec_per_lookup_op"`
    UsecPerMkdirOp                       float64 `json:"usec_per_mkdir_op"`
    UsecPerPathconfOp                    float64 `json:"usec_per_pathconf_op"`
    UsecPerReadOp                        float64 `json:"usec_per_read_op"`
    UsecPerReaddirOp                     float64 `json:"usec_per_readdir_op"`
    UsecPerReaddirplusOp                 float64 `json:"usec_per_readdirplus_op"`
    UsecPerReadlinkOp                    float64 `json:"usec_per_readlink_op"`
    UsecPerRemoveOp                      float64 `json:"usec_per_remove_op"`
    UsecPerRenameOp                      float64 `json:"usec_per_rename_op"`
    UsecPerRmdirOp                       float64 `json:"usec_per_rmdir_op"`
    UsecPerSetattrOp                     float64 `json:"usec_per_setattr_op"`
    UsecPerSymlinkOp                     float64 `json:"usec_per_symlink_op"`
    UsecPerWriteOp                       float64 `json:"usec_per_write_op"`
}

type ArraysNfsPerformanceList struct {
    CntToken       string              `json:"continuation_token"`
    TotalItemCnt   int                 `json:"total_item_count"`
    Items          []ArrayNfsPerformance `json:"items"`
}

func (fb *FBClient) GetArraysNfsPerformance() *ArraysNfsPerformanceList {
    result := new(ArraysNfsPerformanceList)
    _, err := fb.RestClient.R().
                    SetResult(&result).
                    Get("/arrays/nfs-specific-performance")
    
    if err != nil {
        fb.Error = err
    }
    return result
}
