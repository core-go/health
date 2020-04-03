package health

import "context"

type HealthService interface {
	Name() string
	Check(ctx context.Context) error
	Build(ctx context.Context, err error) map[string]interface{}
}
