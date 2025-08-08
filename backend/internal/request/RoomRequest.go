package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterRoomRoutes(g *echo.Group) {
	room := g.Group("/room")
	room.GET("", services.GetAllRoom)
	room.POST("", services.CreateRoom)
	room.PUT("/:id", services.UpdateRoom)
	room.DELETE("/:id", services.DeleteRoom)

	room.POST("/bulk", services.UploadRoomFromExcel)
}
