package client


type BucketDefaults struct {
	HardLimitEnabled      bool         `json:"hard_limit_enabled"`
	QuotaLimit            int          `json:"quota_limit"`
}

type PublicAccessConfig struct {
	BlockNewPublicPolicies      bool         `json:"block_new_public_policies"`
	BlockPublicAccess           bool         `json:"block_public_access"`
}

type ObjectStoreAccount struct {
	Name                       string             `json:"name"`
	Id                         string             `json:"id"`
	Created                    int                `json:"created"`
	ObjectCount                int                `json:"object_count"`
	BucketDefaults             BucketDefaults     `json:"bucket_defaults"`
	HardLimitEnabled           bool               `json:"hard_limit_enabled"`
	Space                      Space              `json:"space"`
	QuotaLimit                 int                `json:"quota_limit"`
	PublicAccessConfig         PublicAccessConfig `json:"public_access_config"`
}

type ObjectStoreAccountsList struct {
	CntToken     string               `json:"continuation_token"`
	TotalItemCnt int                  `json:"total_item_count"`
	Items        []ObjectStoreAccount `json:"items"`
	Total        ObjectStoreAccount   `json:"total"`
}

func (fb *FBClient) GetObjectStoreAccounts() *ObjectStoreAccountsList {
	uri := "/object-store-accounts"
	result := new(ObjectStoreAccountsList)
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
