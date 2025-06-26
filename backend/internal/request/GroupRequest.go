package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterGroupRoutes(g *echo.Group) {
	group := g.Group("/group")
	group.GET("", services.GetAllAssigments)
}
