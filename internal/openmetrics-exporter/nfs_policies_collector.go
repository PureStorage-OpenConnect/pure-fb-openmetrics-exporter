package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type NfsPoliciesCollector struct {
	NfsPolicyDesc *prometheus.Desc
	Client        *client.FBClient
}

func (c *NfsPoliciesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.NfsPolicyDesc
}

func (c *NfsPoliciesCollector) Collect(ch chan<- prometheus.Metric) {
	policies := c.Client.GetNFSExportPolicies()
	if len(policies.Items) > 0 {
		for _, pol := range policies.Items {
			for _, r := range pol.Rules {
				auid := strconv.Itoa(r.AnonUid)
				agid := strconv.Itoa(r.AnonGid)
				sec := strconv.FormatBool(r.Secure)
				f32 := strconv.FormatBool(r.FileId32bit)
				atime := strconv.FormatBool(r.Atime)
				idx := strconv.Itoa(r.Index)
				for _, s := range r.Security {
					ch <- prometheus.MustNewConstMetric(
						c.NfsPolicyDesc,
						prometheus.GaugeValue,
						1.0,
						pol.Name, r.Client, r.Permission, s, r.Access, auid, agid, sec, f32, atime, idx,
					)
				}
			}
		}
	}
}

func NewNfsPoliciesCollector(fb *client.FBClient) *NfsPoliciesCollector {
	return &NfsPoliciesCollector{
		NfsPolicyDesc: prometheus.NewDesc(
			"purefb_nfs_export_rule",
			"FlashBlade NFS export rule",
			[]string{"policy", "client", "permission", "security", "access", "anon_uid", "anon_gid", "secure", "fid32", "atime", "index"},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
