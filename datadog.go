package datadog

import (
	"net/http"
)

var DefaultClient Client

func init() {
	DefaultClient.HTTPClient = http.DefaultClient
}

func UploadTimeSeries(tsReq TimeSeriesRequest) error {
	return DefaultClient.UploadTimeSeries(tsReq)
}
