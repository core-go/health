package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/common-go/health"
)

type GinHealthHandler struct {
	HealthCheckers []health.HealthChecker
}

func NewGinHealthHandler(checkers ...health.HealthChecker) *GinHealthHandler {
	return &GinHealthHandler{checkers}
}

func (c *GinHealthHandler) Check() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := health.Check(ctx.Request.Context(), c.HealthCheckers)
		if result.Status == health.StatusUp {
			ctx.JSON(http.StatusOK, result)
		} else {
			ctx.JSON(http.StatusInternalServerError, result)
		}
	}
}
