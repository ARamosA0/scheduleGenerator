package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterRoomRoutes(g *echo.Group) {
	teacher := g.Group("/room")
	teacher.GET("", services.GetAllRoom)
	teacher.POST("", services.CreateTeacher)
	// teacher.GET("/:id", GetTeacherByID)
	// teacher.PUT("/:id", UpdateTeacher)
	// teacher.DELETE("/:id", DeleteTeacher)
}
