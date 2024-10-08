package client

type Http struct {
	Enabled bool `json:"enabled"`
}

type MultiProtocol struct {
	AccessControlStyle string `json:"access_control_style"`
	SafeguardAcls      bool   `json:"safeguard_acls"`
}

type Nfs struct {
	V3Enabled    bool             `json:"v3_enabled"`
	V41Enabled   bool             `json:"v4_1_enabled"`
	Rules        string           `json:"rules"`
	ExportPolicy ExportRulePolicy `json:"export_policy"`
}

type Smb struct {
	Enabled bool `json:"enabled"`
}

type Location struct {
	Name         string `json:"name"`
	Id           string `json:"id"`
	ResourceType string `jsoon:"resource_type"`
}

type Source struct {
	Name         string   `json:"name"`
	Id           string   `json:"id"`
	ResourceType string   `json:"resource_type"`
	DipslayName  string   `json:"display_name"`
	IsLocal      bool     `json:"is_local"`
	Location     Location `json:"location"`
}

type FileSystem struct {
	Name                       string        `json:"name"`
	Id                         string        `json:"id"`
	Created                    int           `json:"created"`
	DefaultGroupQuota          int           `json:"default_group_quota"`
	DefaultUserQuota           int           `json:"default_user_quota"`
	Destroyed                  bool          `json:"destroyed"`
	FastRemoveDirectoryEnabled bool          `json:"fast_remove_directory_enabled"`
	HardLimitEnabled           bool          `json:"hard_limit_enabled"`
	Http                       Http          `json:"http"`
	MultiProtocol              MultiProtocol `json:"multi_protocol"`
	Nfs                        Nfs           `json:"nfs"`
	PromotionStatus            string        `json:"promotion_status"`
	Provisioned                int           `json:"provisioned"`
	RequestedPromotionState    string        `json:"requested_promotion_state"`
	Smb                        Smb           `json:"smb"`
	SnapshotDirectoryEnabled   bool          `json:"snapshot_directory_enabled"`
	Source                     Source        `json:"source"`
	Space                      Space         `json:"space"`
	TimeRemaining              int           `json:"time_remaining"`
	Writable                   bool          `json:"writable"`
}

type FileSystemsList struct {
	CntToken     string       `json:"continuation_token"`
	TotalItemCnt int          `json:"total_item_count"`
	Items        []FileSystem `json:"items"`
	Total        FileSystem   `json:"total"`
}

func (fb *FBClient) GetFileSystems() *FileSystemsList {
	uri := "/file-systems"
	result := new(FileSystemsList)
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
