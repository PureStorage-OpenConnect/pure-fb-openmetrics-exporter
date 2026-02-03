package collectors

import (
	"testing"
	"unicode"

	"github.com/google/go-cmp/cmp"
	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
)

func normalizeProtoText(s string) string {
	// The proto text output used by Metric.String() can differ slightly between toolchain
	// versions (e.g., extra spaces between fields). We normalize whitespace outside of
	// quoted strings to keep unit tests stable.
	out := make([]rune, 0, len(s))
	inQuotes := false
	spacePending := false
	for _, r := range s {
		if r == '"' {
			if spacePending {
				out = append(out, ' ')
				spacePending = false
			}
			inQuotes = !inQuotes
			out = append(out, r)
			continue
		}
		if inQuotes {
			out = append(out, r)
			continue
		}
		if unicode.IsSpace(r) {
			spacePending = true
			continue
		}
		if spacePending {
			out = append(out, ' ')
			spacePending = false
		}
		out = append(out, r)
	}
	return string(out)
}

func metricsCheck(t *testing.T, c prometheus.Collector, want map[string]bool) {
	chM := make(chan prometheus.Metric)
	go func() {
		c.Collect(chM)
		close(chM)
	}()
	var buff io_prometheus_client.Metric
	metrics := make(map[string]bool)
	for m := range chM {
		m.Write(&buff)
		metrics[normalizeProtoText(buff.String())] = true
	}
	if diff := cmp.Diff(want, metrics); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func metricsCheckWithDesc(t *testing.T, c prometheus.Collector, want map[string]bool) {
	chM := make(chan prometheus.Metric)
	go func() {
		c.Collect(chM)
		close(chM)
	}()

	var buff io_prometheus_client.Metric
	metrics := make(map[string]bool)
	for m := range chM {
		m.Write(&buff)
		metrics[m.Desc().String()+" "+normalizeProtoText(buff.String())] = true
	}

	if len(metrics) != len(want) {
		t.Fatalf("unexpected metric count: got=%d want=%d", len(metrics), len(want))
	}
	for k := range want {
		if !metrics[k] {
			t.Fatalf("missing metric: %s", k)
		}
	}
}
