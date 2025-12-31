package client

import "strings"

const (
	protocolAll = "all"
	protocolNFS = "NFS"
	protocolSMB = "SMB"
)

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
	if protocol != protocolAll && protocol != protocolNFS && protocol != protocolSMB {
		return result
	}
	const chunkSize = 10

	var filesystems []FileSystem
	switch protocol {
	case protocolNFS:
		for _, fs := range f.Items {
			if fs.Nfs.V3Enabled || fs.Nfs.V41Enabled {
				filesystems = append(filesystems, fs)
			}
		}
	case protocolSMB:
		for _, fs := range f.Items {
			if fs.Smb.Enabled {
				filesystems = append(filesystems, fs)
			}
		}
	case protocolAll:
		filesystems = f.Items
	}

	for i := 0; i < len(filesystems); i += chunkSize {
		names := make([]string, 0, chunkSize)
		for _, fs := range filesystems[i:min(i+chunkSize, len(filesystems))] {
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
	return result
}
