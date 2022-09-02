package collectors

import (
    "strconv"
    "purestorage/fb-openmetrics-exporter/internal/rest-client"
    "github.com/prometheus/client_golang/prometheus"
)

type HardwareCollector struct {
    HardwareDesc     *prometheus.Desc
    Client           *client.FBClient
}

func (c *HardwareCollector) Describe(ch chan<- *prometheus.Desc) {
    prometheus.DescribeByCollect(c, ch)
}

func (c *HardwareCollector) Collect(ch chan<- prometheus.Metric) {
    hw := c.Client.GetHardware()
    cstate := 0.0
    for _, comp := range hw.Items {
        switch status := comp.Status; status {
        case "not_installed":
            continue
        case "healthy":
            cstate = 1.0
        case "unused":
            cstate = 2.0
        default:
            cstate = 0.0
        }
        ch <- prometheus.MustNewConstMetric(
                c.HardwareDesc,
                prometheus.GaugeValue,
                cstate,
                comp.Type, comp.Name, strconv.Itoa(comp.Index), strconv.Itoa(comp.Slot),
        )
    }
}

func NewHardwareCollector(fb *client.FBClient) *HardwareCollector {
    return &HardwareCollector{
        HardwareDesc: prometheus.NewDesc(
            "purefb_hardware_health",
            "FlashBlade hardware component health status",
            []string{"type", "name", "index", "slot"},
            prometheus.Labels{},
        ),
        Client: fb,
    }
}
