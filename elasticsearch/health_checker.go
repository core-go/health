package elasticsearch

import (
	"context"
	"time"

	"github.com/elastic/go-elasticsearch"
)

type HealthChecker struct {
	client  *elasticsearch.Client
	name    string
	timeout time.Duration
}

func NewElasticSearchHealthChecker(client *elasticsearch.Client, name string, timeouts ...time.Duration) *HealthChecker {
	var timeout time.Duration
	if len(timeouts) >= 1 {
		timeout = timeouts[0]
	} else {
		timeout = 4 * time.Second
	}
	return &HealthChecker{client, name, timeout}
}

func NewHealthChecker(client *elasticsearch.Client, options ...string) *HealthChecker {
	var name string
	if len(options) > 0 && len(options[0]) > 0 {
		name = options[0]
	} else {
		name = "elasticsearch"
	}
	return NewElasticSearchHealthChecker(client, name, 4 * time.Second)
}

func (e *HealthChecker) Name() string {
	return e.name
}

func (e *HealthChecker) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	_, err := e.client.Ping()
	if err != nil {
		return nil, err
	}
	res["status"] = "success"
	return res, nil
}

func (e *HealthChecker) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
