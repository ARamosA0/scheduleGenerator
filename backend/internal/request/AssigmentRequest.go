package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterAssigmentsRoutes(g *echo.Group) {
	assigment := g.Group("/assigment")
	assigment.GET("", services.GetAllAssigments)
	assigment.POST("", services.CreateAssigment)
	assigment.PUT("/:id", services.UpdateAssigment)
	assigment.DELETE("/:id", services.DeleteAssigment)
}
