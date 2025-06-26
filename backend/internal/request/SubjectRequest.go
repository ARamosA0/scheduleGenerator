package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterSubjectRoutes(g *echo.Group) {
	subject := g.Group("/subject")
	subject.GET("", services.GetAllAssigments)
}
