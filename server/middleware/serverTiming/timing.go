// From https://github.com/p768lwy3/gin-server-timing

package serverTiming

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Timing struct {
	Metrics []*Metric
	sync.Mutex
}

func (h *Timing) NewMetric(name string) *Metric {
	return h.Add(&Metric{Name: name})
}

func (h *Timing) NewMetricWithDesc(name, desc string) *Metric {
	return h.Add(&Metric{Name: name, Desc: desc})
}

func (h *Timing) NewReckonMetric(name string) *Metric {
	return h.Add(&Metric{Name: name}).Start()
}

func (h *Timing) NewReckonMetricWithDesc(name, desc string) *Metric {
	return h.Add(&Metric{Name: name, Desc: desc}).Start()
}

func (h *Timing) Add(m *Metric) *Metric {
	if h == nil {
		return m
	}

	h.Lock()
	defer h.Unlock()
	h.Metrics = append(h.Metrics, m)
	return m
}

func (h *Timing) String() string {
	parts := make([]string, 0, len(h.Metrics))
	for _, m := range h.Metrics {
		parts = append(parts, m.String())
	}

	return strings.Join(parts, metricSplitChar)
}

func headerEncodeParam(key, value string) string {
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return fmt.Sprintf(`%s=%s`, key, value)
	}

	return fmt.Sprintf(`%s=%q`, key, value)
}
