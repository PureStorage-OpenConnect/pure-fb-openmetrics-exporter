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
		temp := new(FileSystemsPerformanceList)
		for i := 0; i < len(f.Items); i += 5 {
			n := ""
			for j := 0; (j < 5) && (i+j < len(f.Items)); j++ {
				n = n + f.Items[i+j].Name + ","
			}
			n = n[:len(n)-1]
			res, _ := fb.RestClient.R().
				SetResult(&temp).
				SetQueryParam("names", n).
				SetQueryParam("protocol", protocol).
				Get(uri)
			if res.StatusCode() == 401 {
				fb.RefreshSession()
				fb.RestClient.R().
					SetResult(&temp).
					SetQueryParam("names", n).
					SetQueryParam("protocol", protocol).
					Get(uri)
			}
			result.Items = append(result.Items, temp.Items...)
		}
	}
	return result
}
