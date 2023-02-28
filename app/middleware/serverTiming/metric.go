// From https://github.com/p768lwy3/gin-server-timing

package serverTiming

import (
	"strconv"
	"strings"
	"time"
)

type Metric struct {
	Name     string
	Duration time.Duration
	Desc     string

	hasDuration bool
	startTime   time.Time
}

func (m *Metric) SetDesc(desc string) *Metric {
	m.Desc = desc
	return m
}

func (m *Metric) Start() *Metric {
	m.startTime = time.Now()
	m.hasDuration = true
	return m
}

func (m *Metric) Stop() *Metric {
	if !m.startTime.IsZero() && m.hasDuration {
		m.Duration = time.Since(m.startTime)
	}

	return m
}

func (m *Metric) String() string {
	parts := make([]string, 1)
	parts[0] = m.Name

	if m.Desc != "" {
		parts = append(parts, headerEncodeParam(paramNameDesc, m.Desc))
	}

	if m.hasDuration {
		parts = append(parts, headerEncodeParam(
			paramNameDur,
			strconv.FormatFloat(float64(m.Duration)/float64(time.Millisecond), 'f', -1, 64),
		))
	}

	return strings.Join(parts, propertiesSplitChat)
}
