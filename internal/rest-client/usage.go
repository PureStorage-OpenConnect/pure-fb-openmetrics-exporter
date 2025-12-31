package client

import "sync"

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FileSystemShort struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type UsageGroups struct {
	Name                   string          `json:"name"`
	FileSystem             FileSystemShort `json:"file_system"`
	FileSystemDefaultQuota float64         `json:"file_system_default_quota"`
	Quota                  float64         `json:"quota"`
	Usage                  float64         `json:"usage"`
	Group                  Group           `json:"group"`
}

type UsageGroupsList struct {
	CntToken     string        `json:"continuation_token"`
	TotalItemCnt int           `json:"total_item_count"`
	Items        []UsageGroups `json:"items"`
}

type UsageUsers struct {
	Name                   string          `json:"name"`
	FileSystem             FileSystemShort `json:"file_system"`
	FileSystemDefaultQuota float64         `json:"file_system_default_quota"`
	Quota                  float64         `json:"quota"`
	Usage                  float64         `json:"usage"`
	User                   User            `json:"user"`
}

type UsageUsersList struct {
	CntToken     string       `json:"continuation_token"`
	TotalItemCnt int          `json:"total_item_count"`
	Items        []UsageUsers `json:"items"`
}

func (fb *FBClient) GetUsageUsers(f *FileSystemsList) *UsageUsersList {
	uri := "/usage/users"
	result := new(UsageUsersList)
	var wg sync.WaitGroup
	outputChan := make(chan *UsageUsersList, len(f.Items))
	for _, fs := range f.Items {
		wg.Add(1)
		go func(fs *FileSystem) {
			defer wg.Done()
			temp := new(UsageUsersList)
			res, _ := fb.RestClient.R().
				SetResult(&temp).
				SetQueryParam("file_system_ids", fs.Id).
				Get(uri)
			if res.StatusCode() == 401 {
				fb.RefreshSession()
				fb.RestClient.R().
					SetResult(&temp).
					SetQueryParam("file_system_ids", fs.Id).
					Get(uri)
			}
			outputChan <- temp
		}(&fs)
	}
	wg.Wait()
	close(outputChan)
	for temp := range outputChan {
		result.Items = append(result.Items, temp.Items...)
	}
	return result
}

func (fb *FBClient) GetUsageGroups(f *FileSystemsList) *UsageGroupsList {
	uri := "/usage/groups"
	result := new(UsageGroupsList)
	var wg sync.WaitGroup
	outputChan := make(chan *UsageGroupsList, len(f.Items))
	for _, fs := range f.Items {
		wg.Add(1)
		go func(fs *FileSystem) {
			defer wg.Done()
			temp := new(UsageGroupsList)
			res, _ := fb.RestClient.R().
				SetResult(&temp).
				SetQueryParam("file_system_ids", fs.Id).
				Get(uri)
			if res.StatusCode() == 401 {
				fb.RefreshSession()
				fb.RestClient.R().
					SetResult(&temp).
					SetQueryParam("file_system_ids", fs.Id).
					Get(uri)
			}
			outputChan <- temp
		}(&fs)
	}
	wg.Wait()
	close(outputChan)
	for temp := range outputChan {
		result.Items = append(result.Items, temp.Items...)
	}
	return result
}
