package client

type Account struct {
	Name         string `json:"name"`
	Id           string `json:"id"`
	ResourceType string `json:"resource_type"`
}

type EradicationConfig struct {
	ManualEradication string `json:"manual_eradication"`
	EradicationDelay  int    `json:"eradication_delay"`
}

type ObjectLockConfig struct {
	Enabled              bool   `json:"enabled"`
	FreezeLockedObjects  bool   `json:"freeze_locked_objects"`
	DefaultRetention     int    `json:"default_retention"`
	DefaultRetentionMode string `json:"default_retention_mode"`
}

type Bucket struct {
	Name             string            `json:"name"`
	Id               string            `json:"id"`
	Account          Account           `json:"account"`
	Created          int               `json:"created"`
	Destroyed        bool              `json:"destroyed"`
	TimeRemaining    int               `json:"time_remaining"`
	ObjectCount      float64           `json:"object_count"`
	Space            Space             `json:"space"`
	Versioning       string            `json:"versioning"`
	BucketType       string            `json:"bucket_type"`
	QuotaLimit       float64           `json:"quota_limit"`
	HardLimitEnabled bool              `json:"hard_limit_enabled"`
	RetentionLock    string            `json:"retention_lock"`
	EradicationCfg   EradicationConfig `json:"eradication_config"`
	ObjectLockCfg    ObjectLockConfig  `json:"object_lock_config"`
}

type BucketsList struct {
	CntToken     string   `json:"continuation_token"`
	TotalItemCnt int      `json:"total_item_count"`
	Items        []Bucket `json:"items"`
	Total        Bucket   `json:"total"`
}

func (fb *FBClient) GetBuckets() *BucketsList {
	uri := "/buckets"
	result := new(BucketsList)
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
