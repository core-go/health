package echo

import (
	"github.com/core-go/health"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthHandler struct {
	Checkers []health.Checker
}

func NewHandler(checkers ...health.Checker) *HealthHandler {
	return &HealthHandler{checkers}
}

func (c *HealthHandler) Check() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result := health.Check(ctx.Request().Context(), c.Checkers)
		if result.Status == health.StatusUp {
			return ctx.JSON(http.StatusOK, result)
		} else {
			return ctx.JSON(http.StatusInternalServerError, result)
		}
	}
}
