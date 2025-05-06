package client

type Reference struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type FixedReference struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type FixedReferenceNameOnly struct {
	Name string `json:"name"`
}

type ReferenceWritable struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type ObjectBacklog struct {
	BytesCount     int64 `json:"bytes_count"`
	DeleteOpsCount int64 `json:"delete_ops_count"`
	OtherOpsCount  int64 `json:"other_ops_count"`
	PutOpsCount    int64 `json:"put_ops_count"`
}

type BucketReplicaLink struct {
	Id                string                 `json:"id"`
	Direction         string                 `json:"direction"`
	Lag               int64                  `json:"lag"`
	StatusDetails     string                 `json:"status_details"`
	Context           Reference              `json:"context"`
	CascadingEnabled  bool                   `json:"cascading_enabled"`
	LocalBucket       FixedReference         `json:"local_bucket"`
	ObjectBacklog     ObjectBacklog          `json:"object_backlog"`
	Paused            bool                   `json:"paused"`
	RecoveryPoint     int64                  `json:"recovery_point"`
	Remote            FixedReference         `json:"remote"`
	RemoteBucket      FixedReferenceNameOnly `json:"remote_bucket"`
	RemoteCredentials ReferenceWritable      `json:"remote_credentials"`
	Status            string                 `json:"status"`
}

type BucketsReplicaLinksList struct {
	CntToken     string              `json:"continuation_token"`
	TotalItemCnt int32               `json:"total_item_count"`
	Items        []BucketReplicaLink `json:"items"`
	Total        BucketReplicaLink   `json:"total"`
}

func (fb *FBClient) GetBucketsReplicaLinks() *BucketsReplicaLinksList {
	uri := "/bucket-replica-links"
	result := new(BucketsReplicaLinksList)
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
