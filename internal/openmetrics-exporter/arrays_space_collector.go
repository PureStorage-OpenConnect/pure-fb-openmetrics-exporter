package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type ArraySpaceCollector struct {
	ReductionDesc   *prometheus.Desc
	SpaceDesc       *prometheus.Desc
	UtilizationDesc *prometheus.Desc
	ParityDesc      *prometheus.Desc
	Client          *client.FBClient
}

func (c *ArraySpaceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *ArraySpaceCollector) Collect(ch chan<- prometheus.Metric) {
	objtypes := []string{"array", "file-system", "object-store"}
	for _, t := range objtypes {
		arrayspace := c.Client.GetArraysSpace(t)

		if len(arrayspace.Items) == 0 {
			continue
		}

		as := arrayspace.Items[0]
		if as.Space.DataReduction != nil {
			ch <- prometheus.MustNewConstMetric(
				c.ReductionDesc,
				prometheus.GaugeValue,
				float64(*as.Space.DataReduction),
				t,
			)
		}
		if as.Space.Snapshots != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.Snapshots),
				t, "snapshots",
			)
		}
		if as.Space.TotalPhysical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.TotalPhysical),
				t, "total_physical",
			)
		}
		if as.Space.Unique != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.Unique),
				t, "unique",
			)
		}
		if as.Space.Virtual != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.Virtual),
				t, "virtual",
			)
		}
		if as.Space.TotalProvisioned != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.TotalProvisioned),
				t, "total_provisioned",
			)
		}
		if as.Space.AvailableProvisioned != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.AvailableProvisioned),
				t, "available_provisioned",
			)
		}
		if as.Space.AvailableRatio != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.AvailableRatio),
				t, "available_ratio",
			)
		}
		if as.Space.Destroyed != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.Destroyed),
				t, "destroyed",
			)
		}
		if as.Space.DestroyedVirtual != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.DestroyedVirtual),
				t, "destroyed_virtual",
			)
		}
		if as.Space.Shared != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*as.Space.Shared),
				t, "shared",
			)
		}
		ch <- prometheus.MustNewConstMetric(
			c.ParityDesc,
			prometheus.GaugeValue,
			as.Parity,
			t,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(as.Capacity),
			t, "capacity",
		)

		if t == "array" {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(as.Capacity-*as.Space.TotalPhysical),
				t, "empty",
			)
			ch <- prometheus.MustNewConstMetric(
				c.UtilizationDesc,
				prometheus.GaugeValue,
				float64(*as.Space.TotalPhysical)/float64(as.Capacity)*100,
				t,
			)
		}
	}
}

func NewArraySpaceCollector(fb *client.FBClient) *ArraySpaceCollector {
	return &ArraySpaceCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefb_array_space_data_reduction_ratio",
			"FlashBlade space data reduction",
			[]string{"type"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefb_array_space_bytes",
			"FlashBlade space in bytes",
			[]string{"type", "space"},
			prometheus.Labels{},
		),
		ParityDesc: prometheus.NewDesc(
			"purefb_array_space_parity",
			"FlashBlade space parity",
			[]string{"type"},
			prometheus.Labels{},
		),
		UtilizationDesc: prometheus.NewDesc(
			"purefb_array_space_utilization",
			"FlashBlade array space utilization in percent",
			[]string{"type"},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
