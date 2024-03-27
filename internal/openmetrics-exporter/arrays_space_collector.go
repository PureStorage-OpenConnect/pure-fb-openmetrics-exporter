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
		ch <- prometheus.MustNewConstMetric(
			c.ReductionDesc,
			prometheus.GaugeValue,
			float64(as.Space.DataReduction),
			t,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(as.Space.Destroyed),
			t, "destroyed",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(as.Space.DestroyedVirtual),
			t, "destroyed_virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(as.Space.Snapshots),
			t, "snapshots",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(as.Space.TotalPhysical),
			t, "total_physical",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(as.Space.Unique),
			t, "unique",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(as.Space.Virtual),
			t, "virtual",
		)
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
				float64(as.Capacity-as.Space.TotalPhysical),
				t, "empty",
			)
			ch <- prometheus.MustNewConstMetric(
				c.UtilizationDesc,
				prometheus.GaugeValue,
				float64(as.Space.TotalPhysical)/float64(as.Capacity)*100,
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
