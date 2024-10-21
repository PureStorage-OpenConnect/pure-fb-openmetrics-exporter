package collectors

import (
	"fmt"
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
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
		al[fmt.Sprintf("%s\n%d\n%s\n%s\n%d\n%s\n%s\n%s",
			strings.Replace(alert.Action, "\n", " ", -1),
			alert.Code,
			alert.ComponentName,
			alert.ComponentType,
			alert.Created,
			alert.KBurl,
			alert.Severity,
			strings.Replace(alert.Summary, "\n", " ", -1))] += 1
	}
	for a, n := range al {
		alert := strings.Split(a, "\n")
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
