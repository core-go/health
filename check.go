package health

import "context"

func Check(ctx context.Context, services []HealthService) Health {
	health := Health{}
	health.Status = StatusUp
	healths := make(map[string]Health)
	for _, service := range services {
		sub := Health{}
		c := context.Background()
		data0, err := service.Check(c)
		if err == nil {
			sub.Status = StatusUp
			if len(data0) > 0 {
				sub.Data = &data0
			}
		} else {
			sub.Status = StatusDown
			health.Status = StatusDown
			data := service.Build(c, data0, err)
			if len(data) > 0 {
				sub.Data = &data
			}
		}
		healths[service.Name()] = sub
	}
	if len(healths) > 0 {
		health.Details = &healths
	}
	return health
}
