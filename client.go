package datadog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const metricEndpointFormat = "https://api.datadoghq.com/api/v1/series?api_key=%s"

type Client struct {
	APIKey string

	HTTPClient *http.Client
}

// UploadTimeSeries .
func (c *Client) UploadTimeSeries(tsReq TimeSeriesRequest) error {
	if len(tsReq.Series) == 0 {
		return nil
	}

	body, err := json.Marshal(tsReq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(metricEndpointFormat, c.APIKey), bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK && resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("invalid API response: %d", resp.StatusCode)
	}

	io.Copy(ioutil.Discard, resp.Body)
	return nil
}
