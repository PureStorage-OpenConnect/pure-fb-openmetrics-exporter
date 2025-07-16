package client

type FileSystemsPerformanceList struct {
	CntToken     string        `json:"continuation_token"`
	TotalItemCnt int           `json:"total_item_count"`
	Items        []Performance `json:"items"`
	Total        []Performance `json:"total"`
}

func (fb *FBClient) GetFileSystemsPerformance(f *FileSystemsList,
	protocol string) *FileSystemsPerformanceList {
	uri := "/file-systems/performance"
	result := new(FileSystemsPerformanceList)
	switch protocol {
	case "all", "NFS", "SMB":
		res, _ := fb.RestClient.R().
			SetResult(&result).
			SetQueryParam("protocol", protocol).
			Get(uri)
		if res.StatusCode() == 401 {
			fb.RefreshSession()
			fb.RestClient.R().
				SetResult(&result).
				SetQueryParam("protocol", protocol).
				Get(uri)
		}
	}
	return result
}
