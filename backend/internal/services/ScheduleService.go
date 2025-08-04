package services

import (
	"net/http"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func GetSchedule(c echo.Context) error {
	assigment_id := c.Param("assigment_id")
	var schedule model.Schedule
	println("ASSIGMENTID", assigment_id)
	if err := db.DB.Where("assignment_id = ?", assigment_id).First(&schedule).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener horario"})
	}

	return c.JSON(http.StatusOK, schedule)
}
