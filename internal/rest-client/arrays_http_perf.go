package client


type ArrayHttpPerformance struct {
    Name                string   `json:"name"`
    Id                  string   `json:"id"`
    OthersPerSec        float64  `json:"others_per_sec"`
    ReadDirsPerSec      float64  `json:"read_dirs_per_sec"`
    ReadFilesPerSec     float64  `json:"read_files_per_sec"`
    WriteDirsPerSec     float64  `json:"write_dirs_per_sec"`
    WriteFilesPerSec    float64  `json:"write_files_per_sec"`
    Time                int      `json:"time"`
    UsecPerOtherOp      float64  `json:"usec_per_other_op"`
    UsecPerReadDirOp    float64  `json:"usec_per_read_dir_op"`
    UsecPerReadFileOp   float64  `json:"usec_per_read_file_op"`
    UsecPerWriteDirOp   float64  `json:"usec_per_write_dir_op"`
    UsecPerWriteFileOp  float64  `json:"usec_per_write_file_op"`
}

type ArraysHttpPerformanceList struct {
    CntToken       string                 `json:"continuation_token"`
    TotalItemCnt   int                    `json:"total_item_count"`
    Items          []ArrayHttpPerformance `json:"items"`
}

func (fb *FBClient) GetArraysHttpPerformance() *ArraysHttpPerformanceList {
    result := new(ArraysHttpPerformanceList)
    _, err := fb.RestClient.R().
                    SetResult(&result).
                    Get("/arrays/http-specific-performance")
    
    if err != nil {
        fb.Error = err
    }
    return result
}
