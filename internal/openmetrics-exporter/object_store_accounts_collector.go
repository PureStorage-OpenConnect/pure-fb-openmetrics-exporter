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
	ch <- c.ReductionDesc
	ch <- c.SpaceDesc
	ch <- c.ObjectCountDesc
}

func (c *ObjectStoreAccountsCollector) Collect(ch chan<- prometheus.Metric) {
	osa := c.Client.GetObjectStoreAccounts()
	if len(osa.Items) == 0 {
		return
	}
	for _, acct := range osa.Items {
		if acct.Space.DataReduction != nil {
			ch <- prometheus.MustNewConstMetric(
				c.ReductionDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.DataReduction),
				acct.Id, acct.Name,
			)
		}
		if acct.Space.Snapshots != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.Snapshots),
				acct.Id, acct.Name, "snapshots",
			)
		}
		if acct.Space.TotalPhysical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.TotalPhysical),
				acct.Id, acct.Name, "total_physical",
			)
		}
		if acct.Space.Unique != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.Unique),
				acct.Id, acct.Name, "unique",
			)
		}
		if acct.Space.Virtual != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.Virtual),
				acct.Id, acct.Name, "virtual",
			)
		}
		if acct.Space.TotalProvisioned != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.TotalProvisioned),
				acct.Id, acct.Name, "total_provisioned",
			)
		}
		if acct.Space.AvailableProvisioned != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.AvailableProvisioned),
				acct.Id, acct.Name, "available_provisioned",
			)
		}
		if acct.Space.AvailableRatio != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.AvailableRatio),
				acct.Id, acct.Name, "available_ratio",
			)
		}
		if acct.Space.Destroyed != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.Destroyed),
				acct.Id, acct.Name, "destroyed",
			)
		}
		if acct.Space.DestroyedVirtual != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.DestroyedVirtual),
				acct.Id, acct.Name, "destroyed_virtual",
			)
		}
		if acct.Space.Shared != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*acct.Space.Shared),
				acct.Id, acct.Name, "shared",
			)
		}
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
