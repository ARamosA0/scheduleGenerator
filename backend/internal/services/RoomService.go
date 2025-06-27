package services

import (
	"net/http"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func GetAllRoom(c echo.Context) error {
	var room []model.Room

	if err := db.DB.Preload("RoomType").Find(&room).Error; err != nil {
		println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al obtener salas",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, room)
}
