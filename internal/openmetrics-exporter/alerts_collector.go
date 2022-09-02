package collectors

import (
    "purestorage/fb-openmetrics-exporter/internal/rest-client"
    "github.com/prometheus/client_golang/prometheus"
)

type AlertsCollector struct {
    AlertsDesc       *prometheus.Desc
    Client           *client.FBClient
}

func (c *AlertsCollector) Describe(ch chan<- *prometheus.Desc) {
    prometheus.DescribeByCollect(c, ch)
}

func (c *AlertsCollector) Collect(ch chan<- prometheus.Metric) {
    alerts := c.Client.GetAlerts("state='open'")
    if len(alerts.Items) == 0 {
        return 
    }
    for _, alert := range alerts.Items {
        ch <- prometheus.MustNewConstMetric(
                c.AlertsDesc,
                prometheus.GaugeValue,
                1.0,
                alert.Severity, alert.ComponentType, alert.ComponentName,
        )
    }
}

func NewAlertsCollector(fb *client.FBClient) *AlertsCollector {
    return &AlertsCollector{
        AlertsDesc: prometheus.NewDesc(
            "purefb_alerts_open",
            "FlashBlade open alert events",
            []string{"severity", "component_type", "component_name"},
            prometheus.Labels{},
        ),
        Client: fb,
    }
}
