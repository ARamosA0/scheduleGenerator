package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterTeacherRoutes(g *echo.Group) {
	teacher := g.Group("/teachers")
	teacher.GET("", services.GetAllTeachers)
	teacher.POST("", services.CreateTeacher)
	// teacher.GET("/:id", GetTeacherByID)
	// teacher.PUT("/:id", UpdateTeacher)
	// teacher.DELETE("/:id", DeleteTeacher)
}
