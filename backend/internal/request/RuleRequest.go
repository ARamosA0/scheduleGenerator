package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterRuleRoutes(g *echo.Group) {
	rule := g.Group("/rule")
	rule.GET("", services.GetAllAssigments)
}
