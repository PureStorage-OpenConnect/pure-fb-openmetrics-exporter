package client

import "strings"

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
	const chunkSize = 10

	switch protocol {
	case "all", "NFS", "SMB":
		for i := 0; i < len(f.Items); i += chunkSize {
			names := make([]string, 0, chunkSize)
			for _, fs := range f.Items[i:min(i+chunkSize, len(f.Items))] {
				names = append(names, fs.Name)
			}
			n := strings.Join(names, ",")
			temp := new(FileSystemsPerformanceList)
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
