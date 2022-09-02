package collectors

import (
    "purestorage/fb-openmetrics-exporter/internal/rest-client"
    "github.com/prometheus/client_golang/prometheus"
)

type ArraysCollector struct {
    ArraysDesc  *prometheus.Desc
    Client      *client.FBClient
}

func (c *ArraysCollector) Describe(ch chan<- *prometheus.Desc) {
    ch <- c.ArraysDesc
}

func (c *ArraysCollector) Collect(ch chan<- prometheus.Metric) {
    arrays := c.Client.GetArrays()
    if len(arrays.Items) == 0 {
        return
    }
    array := arrays.Items[0]

    ch <- prometheus.MustNewConstMetric(
            c.ArraysDesc,
            prometheus.GaugeValue,
            1.0,
            array.Name, array.Id, array.Os, array.Version,
    )
}

func NewArraysCollector(fb *client.FBClient) *ArraysCollector {
    return &ArraysCollector{
        ArraysDesc: prometheus.NewDesc(
            "purefb_info",
            "FlashBlade system information",
            []string{"array_name", "system_id", "os", "version"},
            prometheus.Labels{},
        ),
        Client: fb,
    }
}
