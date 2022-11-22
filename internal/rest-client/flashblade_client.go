package client

import (
//	"log"
	"errors"
	"crypto/tls"
	"github.com/go-resty/resty/v2"
)

type Client interface {
	GetAlerts(filter string) *AlertsList
	GetArrays() *ArraysList
	GetArraysHttpPerformance() *ArraysHttpPerformanceList
	GetArraysNfsPerformance() *ArraysNfsPerformanceList
	GetArraysPerformance(protocol string) *ArraysPerformanceList
	GetArraysPerformanceReplication() *ArraysPerformanceReplicationList
	GetArraysS3Performance() *ArraysS3PerformanceList
	GetArraysSpace(t string) *ArraysSpaceList
	GetBlades() *BladesList
	GetBuckets() *BucketsList
	GetBucketsPerformance(b *BucketsList) *BucketsPerformanceList
	GetBucketsS3Performance(b *BucketsList) *BucketsS3PerformanceList
	GetClientsPerformance() *ClientsPerformanceList
	GetFileSystems() *FileSystemsList
	GetFileSystemsPerformance(f *FileSystemsList, protocol string) *FileSystemsPerformanceList
	GetHwConnectorsPerformance() *HwConnectorsPerformanceList
	GetHardware() *HardwareList
	GetUsageUsers(f* FileSystemsList) *UsageUsersList
	GetUsageGroups(f* FileSystemsList) *UsageGroupsList
}

type FBClient struct {
	EndPoint   string
	ApiToken   string
	RestClient *resty.Client
	ApiVersion string
	XAuthToken string
	Error      error
}

func NewRestClient(endpoint string, apitoken string, apiversion string, debug bool) *FBClient {
	type ApiVersions struct {
		Versions []string `json:"versions"`
	}
	fb := &FBClient{
		EndPoint:   endpoint,
		ApiToken:   apitoken,
		RestClient: resty.New(),
		XAuthToken: "",
	}
	fb.RestClient.SetBaseURL("https://" + endpoint + "/api")
	fb.RestClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	fb.RestClient.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})

	if debug {
		fb.RestClient.SetDebug(true)
	}
//	fb.RestClient.OnRequestLog(func(rl *resty.RequestLog) error {
//		fmt.Fprintln(os.Stderr, rl)
//		return nil
//	})

	result := new(ApiVersions)
	res, err := fb.RestClient.R().
		SetResult(&result).
		Get("/api_version")
	if err != nil {
		fb.Error = err
		return fb
	}
	if res.StatusCode() != 200 {
		fb.Error = errors.New("Not a valid FlashBlade REST API server")
		return fb
	}
	if len(result.Versions) == 0 {
		fb.Error = errors.New("Not a valid FlashBlade REST API version")
		return fb
	}
	if apiversion == "latest" {
		fb.ApiVersion = result.Versions[len(result.Versions)-1]
	} else {
		fb.ApiVersion = apiversion
	}
	res, err = fb.RestClient.R().
		SetHeader("api-token", apitoken).
		Post("/login")
	if err != nil {
		fb.Error = err
		return fb
	}
	fb.XAuthToken = res.Header().Get("x-auth-token")
	fb.RestClient.SetBaseURL("https://" + endpoint + "/api/" + fb.ApiVersion)
	fb.RestClient.SetHeader("x-auth-token", fb.XAuthToken)
	return fb
}

func (fb *FBClient) Close() *FBClient {
	if fb.XAuthToken == "" {
		return fb
	}
	_, err := fb.RestClient.R().
		SetHeader("x-auth-token", fb.XAuthToken).
		Post("/logout")
	if err != nil {
		fb.Error = err
	}
	return fb
}

func (fb *FBClient) RefreshSession() *FBClient {
	res, err := fb.RestClient.R().
		SetHeader("api-token", fb.ApiToken).
		Post("/login")
	if err != nil {
		fb.Error = err
		return fb
	}
	fb.XAuthToken = res.Header().Get("x-auth-token")
	fb.RestClient.SetHeader("x-auth-token", fb.XAuthToken)
	return fb
}
