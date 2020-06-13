package health

import "context"

type HealthService interface {
	Name() string
	Check(ctx context.Context) (map[string]interface{}, error)
	Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{}
}
