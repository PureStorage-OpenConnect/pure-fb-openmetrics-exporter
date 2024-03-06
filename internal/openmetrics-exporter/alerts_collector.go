package collectors

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fb-openmetrics-exporter/internal/rest-client"
)

type AlertsCollector struct {
	AlertsDesc *prometheus.Desc
	Client     *client.FBClient
}

func (c *AlertsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *AlertsCollector) Collect(ch chan<- prometheus.Metric) {
	alerts := c.Client.GetAlerts("state='open'")
	if len(alerts.Items) == 0 {
		return
	}
	al := make(map[string]float64)
	for _, alert := range alerts.Items {
		al[fmt.Sprintf("%s,%d,%s,%s,%d,%s,%s,%s,%s",
			alert.Action,
			alert.Code,
			alert.ComponentName,
			alert.ComponentType,
			alert.Created,
			alert.KBurl,
			alert.Severity,
			alert.Summary,
		)] += 1
	}
	for a, n := range al { 
		alert := strings.Split(a, ",")
		ch <- prometheus.MustNewConstMetric(
			c.AlertsDesc,
			prometheus.GaugeValue,
			n,
			alert[0],
			alert[1],
			alert[2],
			alert[3],
			alert[4],
			alert[5],
			alert[6],
			alert[7],
		)
	}
}

func NewAlertsCollector(fb *client.FBClient) *AlertsCollector {
	return &AlertsCollector{
		AlertsDesc: prometheus.NewDesc(
			"purefb_alerts_open",
			"FlashBlade open alert events",
			[]string{"action", "code", "component_name", "component_type", "created", "kburl", "severity", "summary"},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
