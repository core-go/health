package v3

import (
	"github.com/core-go/health"
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct {
	Checkers []health.Checker
}

func NewHandler(checkers ...health.Checker) *Handler {
	return &Handler{checkers}
}

func (c *Handler) Check(ctx echo.Context) error {
	result := health.Check(ctx.Request().Context(), c.Checkers)
	if result.Status == health.StatusUp {
		return ctx.JSON(http.StatusOK, result)
	} else {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
}
