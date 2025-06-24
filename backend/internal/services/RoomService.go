package services

import (
	"net/http"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

var room = []model.Room{
	// {
	// 	Name:           "Teacher1",
	// 	AvailableTimes: []time.Time{time.Now()},
	// },
	// {
	// 	Name:           "Teacher2",
	// 	AvailableTimes: []time.Time{time.Now()},
	// },
}

func GetRoom(c echo.Context) error {
	var room []model.Room

	if err := db.DB.Find(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener profesores"})
	}

	return c.JSON(http.StatusOK, room)
}
