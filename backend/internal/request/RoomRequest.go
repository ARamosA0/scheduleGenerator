package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterRoomRoutes(g *echo.Group) {
	teacher := g.Group("/room")
	teacher.GET("", services.GetAllRoom)
	teacher.POST("", services.CreateRoom)
	teacher.PUT("/:id", services.UpdateRoom)
	teacher.DELETE("/:id", services.DeleteRoom)
}
