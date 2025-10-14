package request

import (
	"backend.com/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func HandleServerEvents(g *echo.Group) {
	events := g.Group("/events")
	events.GET("", services.HandleServerEvents)
}
