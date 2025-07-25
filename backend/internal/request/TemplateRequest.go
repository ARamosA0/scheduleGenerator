package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterTemplatesRoutes(g *echo.Group) {
	teacher := g.Group("/template")
	teacher.GET("", services.GetAllTemplates)
	teacher.GET("/:id", services.GetTemplate)
	teacher.POST("", services.CreateTemplate)
	teacher.PUT("/:id", services.UpdateTemplate)
	teacher.DELETE("/:id", services.DeleteTemplate)
}
