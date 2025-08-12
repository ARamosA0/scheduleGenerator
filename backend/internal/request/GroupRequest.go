package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterGroupRoutes(g *echo.Group) {
	group := g.Group("/group")
	group.GET("", services.GetAllGroups)
	group.POST("", services.CreateGroup)
	group.PUT("/:id", services.UpdateGroup)
	group.DELETE("/:id", services.DeleteGroup)

	group.POST("/bulk", services.UploadGroupFromExcel)
}
