package client

type ArrayHttpPerformance struct {
	Name               string  `json:"name"`
	Id                 string  `json:"id"`
	OthersPerSec       float64 `json:"others_per_sec"`
	ReadDirsPerSec     float64 `json:"read_dirs_per_sec"`
	ReadFilesPerSec    float64 `json:"read_files_per_sec"`
	WriteDirsPerSec    float64 `json:"write_dirs_per_sec"`
	WriteFilesPerSec   float64 `json:"write_files_per_sec"`
	Time               int     `json:"time"`
	UsecPerOtherOp     float64 `json:"usec_per_other_op"`
	UsecPerReadDirOp   float64 `json:"usec_per_read_dir_op"`
	UsecPerReadFileOp  float64 `json:"usec_per_read_file_op"`
	UsecPerWriteDirOp  float64 `json:"usec_per_write_dir_op"`
	UsecPerWriteFileOp float64 `json:"usec_per_write_file_op"`
}

type ArraysHttpPerformanceList struct {
	CntToken     string                 `json:"continuation_token"`
	TotalItemCnt int                    `json:"total_item_count"`
	Items        []ArrayHttpPerformance `json:"items"`
}

func (fb *FBClient) GetArraysHttpPerformance() *ArraysHttpPerformanceList {
	uri := "/arrays/http-specific-performance"
	result := new(ArraysHttpPerformanceList)
	res, _ := fb.RestClient.R().
		SetResult(&result).
		Get(uri)
	if res.StatusCode() == 401 {
		fb.RefreshSession()
		fb.RestClient.R().
			SetResult(&result).
			Get(uri)
	}
	return result
}
