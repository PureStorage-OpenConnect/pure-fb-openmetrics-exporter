package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type ObjectStoreAccountsCollector struct {
	ReductionDesc   *prometheus.Desc
	SpaceDesc       *prometheus.Desc
	ObjectCountDesc *prometheus.Desc
	Client          *client.FBClient
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
			acct.Id, acct.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.Snapshots),
			acct.Id, acct.Name, "snapshots",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.TotalPhysical),
			acct.Id, acct.Name, "total_physical",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.Unique),
			acct.Id, acct.Name, "unique",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.Virtual),
			acct.Id, acct.Name, "virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.TotalProvisioned),
			acct.Id, acct.Name, "total_provisioned",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.AvailableProvisioned),
			acct.Id, acct.Name, "available_provisioned",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.Destroyed),
			acct.Id, acct.Name, "destroyed",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(acct.Space.DestroyedVirtual),
			acct.Id, acct.Name, "destroyed_virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ObjectCountDesc,
			prometheus.GaugeValue,
			float64(acct.ObjectCount),
			acct.Id, acct.Name,
		)
	}
}

func NewObjectStoreAccountsCollector(fb *client.FBClient) *ObjectStoreAccountsCollector {
	return &ObjectStoreAccountsCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefb_object_store_accounts_data_reduction_ratio",
			"FlashBlade object store accounts data reduction",
			[]string{"id", "name"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefb_object_store_accounts_space_bytes",
			"FlashBlade object store accounts space in bytes",
			[]string{"id", "name", "space"},
			prometheus.Labels{},
		),
		ObjectCountDesc: prometheus.NewDesc(
			"purefb_object_store_accounts_object_count",
			"FlashBlade object store accounts object count",
			[]string{"id", "name"},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
