package client

type ExportRulePolicy struct {
	Name             string    `json:"name"`
	Id               string    `json:"id"`
        ResourceType     string    `json:"resource_type"`
}

type NFSExportRule struct {
	Name             string            `json:"name"`
	Id               string            `json:"id"`
	Policy           ExportRulePolicy  `json:"policy"`
	Index            int               `json:"index"`
	PolicyVersion    string            `json:"policy_version"`
	Access           string            `json:"access"`
	AnonGid          int               `json:"anongid"`
	AnonUid          int               `json:"anonuid"`
        Atime            bool              `json:"atime"`
	Client           string            `json:"client"`
        FileId32bit      bool              `json:"fileid_32bit"`
	Permission       string            `json:"permission"`
        Secure           bool              `json:"secure"`
        Security         []string          `json:"security"`
}

type NFSExportPolicy struct {
	Name             string            `json:"name"`
	Id               string            `json:"id"`
        Enabled          bool              `json:"enabled"`
        IsLocal          bool              `json:"is_local"`
        Location         Location          `json:"location"`
        Version          string            `json:"version"`
        Rules            []NFSExportRule   `json:"rules"`
        PolicyType       string            `json:"policy_type"`
}

type NFSExportPolicyList struct {
	CntToken     string             `json:"continuation_token"`
	TotalItemCnt int                `json:"total_item_count"`
	Items        []NFSExportPolicy  `json:"items"`
}

func (fb *FBClient) GetNFSExportPolicies() *NFSExportPolicyList {
        uri := "/nfs-export-policies"
	result := new(NFSExportPolicyList)
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
