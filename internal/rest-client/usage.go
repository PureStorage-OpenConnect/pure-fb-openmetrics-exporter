package restclient

type Group struct {
    Id      int      `json:"id"`
    Name    string   `json:"name"`
}

type User struct {
    Id      int     `json:"id"`
    Name    string   `json:"name"`
}

type FileSystemShort struct {
    Id            string   `json:"id"`
    Name          string   `json:"name"`
    ResourceType  string   `json:"resource_type"`
}

type UsageGroups struct {
    Name                     string           `json:"name"`
    FileSystem               FileSystemShort  `json:"file_system"`
    FileSystemDefaultQuota   float64          `json:"file_system_default_quota"`
    Quota                    float64          `json:"quota"`
    Usage                    float64          `json:"usage"`
    Group                    Group            `json:"group"`
}

type UsageGroupsList struct {
    CntToken       string         `json:"continuation_token"`
    TotalItemCnt   int            `json:"total_item_count"`
    Items          []UsageGroups  `json:"items"`
}

type UsageUsers struct {
    Name                     string           `json:"name"`
    FileSystem               FileSystemShort  `json:"file_system"`
    FileSystemDefaultQuota   float64          `json:"file_system_default_quota"`
    Quota                    float64          `json:"quota"`
    Usage                    float64          `json:"usage"`
    User                     User             `json:"user"`
}

type UsageUsersList struct {
    CntToken       string        `json:"continuation_token"`
    TotalItemCnt   int           `json:"total_item_count"`
    Items          []UsageUsers  `json:"items"`
}

func (fb *FBClient) GetUsageUsers(f *FileSystemsList) *UsageUsersList {
    result := new(UsageUsersList)
    temp := new(UsageUsersList)
    for i := 0; i < len(f.Items); i++ {
        _, err := fb.RestClient.R().
                        SetResult(&temp).
                        SetQueryParam("file_system_ids", f.Items[i].Id).
                        Get("/usage/users")
    
        if err != nil {
            fb.Error = err
        }
        result.Items = append(result.Items, temp.Items...)
    }
    return result
}

func (fb *FBClient) GetUsageGroups(f *FileSystemsList) *UsageGroupsList {
    result := new(UsageGroupsList)
    temp := new(UsageGroupsList)
    for i := 0; i < len(f.Items); i++ {
        _, err := fb.RestClient.R().
                        SetResult(&temp).
                        SetQueryParam("file_system_ids", f.Items[i].Id).
                        Get("/usage/groups")
    
        if err != nil {
            fb.Error = err
        }
        result.Items = append(result.Items, temp.Items...)
    }
    return result
}
