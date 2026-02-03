package client

type RemoteBucket struct {
	Name string `json:"name"`
}

type BucketReplicaLink struct {
	Id               string        `json:"id"`
	Direction        string        `json:"direction"`
	Lag              float64       `json:"lag"`
	StatusDetails    string        `json:"status_details"`
	Context          *ResourceRef  `json:"context,omitempty"`
	CascadingEnabled bool          `json:"cascading_enabled"`
	LocalBucket      *ResourceRef  `json:"local_bucket,omitempty"`
	ObjectBacklog    *ObjectBacklog `json:"object_backlog,omitempty"`
	Paused           bool          `json:"paused"`
	RecoveryPoint    int64         `json:"recovery_point"`
	Remote           *ResourceRef  `json:"remote,omitempty"`
	RemoteBucket     *RemoteBucket `json:"remote_bucket,omitempty"`
	RemoteCredentials *ResourceRef `json:"remote_credentials,omitempty"`
	Status           string        `json:"status"`
}

type BucketReplicaLinksList struct {
	CntToken     string             `json:"continuation_token"`
	TotalItemCnt int                `json:"total_item_count"`
	Items        []BucketReplicaLink `json:"items"`
	Total        *BucketReplicaLink  `json:"total,omitempty"`
}

func (fb *FBClient) GetBucketReplicaLinks() *BucketReplicaLinksList {
	uri := "/bucket-replica-links"
	result := new(BucketReplicaLinksList)
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


