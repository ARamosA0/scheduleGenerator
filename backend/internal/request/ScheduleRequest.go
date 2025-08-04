package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterScheduleRoutes(g *echo.Group) {
	teacher := g.Group("/schedule")
	teacher.GET("/:assigment_id", services.GetSchedule)
}
