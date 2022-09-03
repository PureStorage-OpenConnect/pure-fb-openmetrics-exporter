package collectors

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fb-openmetrics-exporter/internal/rest-client"
)

func Collector(ctx context.Context, endpoint string, apitoken string, apiver string, metrics string, registry *prometheus.Registry) bool {
	fbclient := client.NewRestClient(endpoint, apitoken, apiver)
	filesystems := fbclient.GetFileSystems()
	buckets := fbclient.GetBuckets()
	defer fbclient.Close()

	registry.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)
	if metrics == "all" || metrics == "array" {
		arrayCollector := NewArraysCollector(fbclient)
		perfCollector := NewPerfCollector(fbclient)
		s3perfCollector := NewS3PerfCollector(fbclient)
		httpPerfCollector := NewHttpPerfCollector(fbclient)
		nfsPerfCollector := NewNfsPerfCollector(fbclient)
		perfReplCollector := NewPerfReplicationCollector(fbclient)
		bucketsPerfCollector := NewBucketsPerfCollector(fbclient, buckets)
		buckestS3PerfCollector := NewBucketsS3PerfCollector(fbclient, buckets)
		filesystemsPerfCollector := NewFileSystemsPerfCollector(fbclient, filesystems)
		filesystemsSpaceCollector := NewFileSystemsSpaceCollector(filesystems)
		arraySpaceCollector := NewArraySpaceCollector(fbclient)
		bucketsSpaceCollector := NewBucketsSpaceCollector(buckets)
		alertsCollector := NewAlertsCollector(fbclient)
		hardwareCollector := NewHardwareCollector(fbclient)
		hwPerfConnectorsCollector := NewHwConnectorsPerfCollector(fbclient)
		registry.MustRegister(
			arrayCollector,
			perfCollector,
			s3perfCollector,
			httpPerfCollector,
			nfsPerfCollector,
			perfReplCollector,
			bucketsPerfCollector,
			buckestS3PerfCollector,
			filesystemsPerfCollector,
			arraySpaceCollector,
			bucketsSpaceCollector,
			filesystemsSpaceCollector,
			alertsCollector,
			hardwareCollector,
			hwPerfConnectorsCollector,
		)
	}
	if metrics == "all" || metrics == "clients" {
		clientsPerfCollector := NewClientsPerfCollector(fbclient)
		registry.MustRegister(clientsPerfCollector)
	}
	if metrics == "all" || metrics == "usage" {
		usageCollector := NewUsageCollector(fbclient, filesystems)
		registry.MustRegister(usageCollector)
	}
	return true
}
