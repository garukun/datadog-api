package datadog_test

import (
	"time"

	"github.com/garukun/datadog-api"
)

func Example() {
	client := &datadog.Client{
		APIKey: "datadog-api-key",
	}

	tsReq := datadog.TimeSeriesRequest{
		Series: []datadog.TimeSeries{
			{
				Name: "example.metric.gauge",
				Points: []datadog.DataPoint{
					{Timestamp: time.Now().Unix(), Value: 3.14},
					{Timestamp: time.Now().Unix(), Value: 3.16},
				},
			},
		},
	}

	client.UploadTimeSeries(tsReq) // also error handling
}
