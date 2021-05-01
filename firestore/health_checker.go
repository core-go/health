package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
)

type HealthChecker struct {
	name      string
	projectId string
	opts      []option.ClientOption
}

func NewFirestoreHealthChecker(name string, projectId string, opts ...option.ClientOption) *HealthChecker {
	return &HealthChecker{projectId: projectId, name: name, opts: opts}
}

func NewHealthChecker(projectId string, opts ...option.ClientOption) *HealthChecker {
	return NewFirestoreHealthChecker("firestore", projectId, opts...)
}
func (s HealthChecker) Name() string {
	return s.name
}

func (s HealthChecker) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	client, err := firestore.NewClient(ctx, s.projectId, s.opts...)
	if err != nil {
		return res, err
	}
	defer client.Close()
	return res, nil
}

func (s *HealthChecker) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
