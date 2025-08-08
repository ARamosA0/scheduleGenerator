package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterValidateDataRoutes(g *echo.Group) {
	teacher := g.Group("/validate")
	teacher.POST("", services.ValidateFile)
}
