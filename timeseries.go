package datadog

import (
	"bytes"
	"fmt"
	"strconv"
)

// TimeSeriesRequest wraps a list of time series data according the Datadog API.
type TimeSeriesRequest struct {
	Series []TimeSeries `json:"series"`
}

type TimeSeries struct {
	Name   string      `json:"metric"`
	Points []DataPoint `json:"points"`

	Type     Type     `json:"type,omitempty"`
	Interval int      `json:"interval,omitempty"`
	Host     string   `json:"host,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

type Type int

const (
	TypeGauge Type = iota
	TypeRate
	TypeCount
)

func (t Type) MarshalJSON() ([]byte, error) {
	switch t {
	case TypeGauge:
		return []byte(`"gauge"`), nil
	case TypeRate:
		return []byte(`"rate"`), nil
	case TypeCount:
		return []byte(`"count"`), nil
	}

	return nil, fmt.Errorf("invalid metric type: %d", t)
}

type DataPoint struct {
	Timestamp int64   `json:"-"`
	Value     float64 `json:"-"`
}

func (p DataPoint) MarshalJSON() ([]byte, error) {
	if p.Timestamp <= 0 {
		return nil, fmt.Errorf("invalid timestamp: %d", p.Timestamp)
	}

	const (
		roughDataPointSize       = 32
		apiDataPointValueBitSize = 32
	)

	buf := bytes.NewBuffer(make([]byte, 0, roughDataPointSize))
	buf.WriteByte('[')
	buf.WriteString(strconv.FormatInt(p.Timestamp, 10))
	buf.WriteByte(',')
	buf.WriteString(strconv.FormatFloat(p.Value, 'f', 6, apiDataPointValueBitSize))
	buf.WriteByte(']')

	return buf.Bytes(), nil
}
