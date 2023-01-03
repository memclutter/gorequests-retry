package gorequests_retry

import (
	"github.com/hashicorp/go-retryablehttp"
	"net/http"
	"time"
)

type Retry struct {
	RetryMax     int
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration
}

func (m *Retry) ClientOverride(c *http.Client) (*http.Client, error) {
	rc := retryablehttp.NewClient()
	rc.HTTPClient = c
	rc.RetryMax = m.RetryMax
	rc.RetryWaitMin = m.RetryWaitMin
	rc.RetryWaitMax = m.RetryWaitMax
	rc.Logger = nil
	return rc.StandardClient(), nil
}
