package restclient


type FileSystemsPerformanceList struct {
    CntToken       string          `json:"continuation_token"`
    TotalItemCnt   int             `json:"total_item_count"`
    Items          []Performance   `json:"items"`
    Total          []Performance   `json:"total"`
}

func (fb *FBClient) GetFileSystemsPerformance(f *FileSystemsList, 
                                              protocol string) *FileSystemsPerformanceList {
    result := new(FileSystemsPerformanceList)
    switch protocol {
        case "all", "HTTP", "NFS", "SMB", "S3":
            temp := new(FileSystemsPerformanceList)
            for i := 0; i < len(f.Items); i += 5 {
                n := ""
                for j := 0; (j < 5) && (i + j < len(f.Items)); j++ {
                   n = n + f.Items[i+j].Name + "," 
                }
                n = n[:len(n)-1]
                _, err := fb.RestClient.R().
                                SetResult(&temp).
                                SetQueryParam("names", n).
                                SetQueryParam("protocol", protocol).
                                Get("/file-systems/performance")
    
                if err != nil {
                    fb.Error = err
                }
                result.Items = append(result.Items, temp.Items...)
            }
    }
    return result
}
