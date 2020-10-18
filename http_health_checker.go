package health

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type HttpHealthChecker struct {
	name    string
	url     string
	timeout time.Duration
}

func NewHttpHealthChecker(name, url string, timeout time.Duration) *HttpHealthChecker {
	return &HttpHealthChecker{name, url, timeout}
}

func NewDefaultHttpHealthChecker(name, url string) *HttpHealthChecker {
	return &HttpHealthChecker{name, url, 5 * time.Second}
}

func (s *HttpHealthChecker) Name() string {
	return s.name
}

func (s *HttpHealthChecker) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	client := http.Client{
		Timeout: s.timeout,
		// never follow redirects
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(s.url)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return res, fmt.Errorf("time out: %w", e)
	} else if err != nil {
		return res, err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if resp.StatusCode >= 500 {
		return res, fmt.Errorf("status code is: %d", resp.StatusCode)
	}
	return res, nil
}

func (s *HttpHealthChecker) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
