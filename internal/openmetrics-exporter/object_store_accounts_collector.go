package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type ObjectStoreAccountsCollector struct {
	ReductionDesc         *prometheus.Desc
	SpaceDesc             *prometheus.Desc
	ObjectCountDesc       *prometheus.Desc
	Client                *client.FBClient
}

func (c *ObjectStoreAccountsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *ObjectStoreAccountsCollector) Collect(ch chan<- prometheus.Metric) {
	osa := c.Client.GetObjectStoreAccounts()
	if len(osa.Items) == 0 {
		return
	}
	for _, acct := range osa.Items {
		ch <- prometheus.MustNewConstMetric(
			c.ReductionDesc,
			prometheus.GaugeValue,
			float64(acct.Space.DataReduction),
			acct.Name, acct.Id,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.Snapshots),
			acct.Name, acct.Id, "snapshots",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.TotalPhysical),
			acct.Name, acct.Id, "total_physical",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.Unique),
			acct.Name, acct.Id, "unique",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.Virtual),
			acct.Name, acct.Id, "virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.TotalProvisioned),
			acct.Name, acct.Id, "total_provisioned",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.AvailableProvisioned),
			acct.Name, acct.Id, "available_provisioned",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.Destroyed),
			acct.Name, acct.Id, "destroyed",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.DestroyedVirtual),
			acct.Name, acct.Id, "destroyed_virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ObjectCountDesc,
			prometheus.GaugeValue,
			float64(acct.ObjectCount),
			acct.Name, acct.Id,
		)
	}
}

func NewObjectStoreAccountsCollector(fb *client.FBClient) *ObjectStoreAccountsCollector {
	return &ObjectStoreAccountsCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefb_object_store_accounts_data_reduction_ratio",
			"FlashBlade object store accounts data reduction",
			[]string{"name", "id"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefb_object_store_accounts_space_bytes",
			"FlashBlade object store accounts space in bytes",
			[]string{"name", "id", "space"},
			prometheus.Labels{},
		),
		ObjectCountDesc: prometheus.NewDesc(
			"purefb_object_store_accounts_object_count",
			"FlashBlade object store accounts object count",
			[]string{"name", "id"},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
