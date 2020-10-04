package health

import (
	"context"
	"encoding/json"
	"net/http"
)

type HealthHandler struct {
	HealthServices []HealthService
}

func NewHealthHandler(healthServices []HealthService) *HealthHandler {
	return &HealthHandler{healthServices}
}

func (c *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	h := Check(ctx, c.HealthServices)
	bytes, err := json.Marshal(h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if h.Status == StatusDown {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(bytes)
}
