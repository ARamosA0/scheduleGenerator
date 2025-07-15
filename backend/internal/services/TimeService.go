package services

import (
	"net/http"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func GetAllTemplates(c echo.Context) error {
	var template []model.Template

	if err := db.DB.Preload("RoomType").Find(&template).Error; err != nil {
		println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al obtener salas",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, template)
}

func GetTemplate(c echo.Context) error {
	var room model.Room

	if err := db.DB.Find(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener profesores"})
	}

	return c.JSON(http.StatusOK, room)
}

func CreateTemplate(c echo.Context) error {
	var t model.Template

	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	// if !json.Valid(t.DaysRange) {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Invalid daysRange JSON")
	// }

	if err := db.DB.Create(&t).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create template")
	}

	return c.JSON(http.StatusCreated, t)
}

func UpdateTemplate(c echo.Context) error {
	id := c.Param("id")
	var room model.Room
	if err := db.DB.First(&room, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Profesor no encontrado"})
	}

	var input model.Room
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	room.Code = input.Code
	room.Name = input.Name
	room.Capacity = input.Capacity
	room.RoomTypeID = input.RoomTypeID
	room.Floor = input.Floor
	room.Building = input.Building
	room.Observations = input.Observations

	if err := db.DB.Save(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo actualizar"})
	}

	return c.JSON(http.StatusCreated, room)
}

func DeleteTemplate(c echo.Context) error {
	id := c.Param("id")

	var room model.Room
	if err := db.DB.First(&room, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Profesor no encontrado",
		})
	}

	if err := db.DB.Delete(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar el profesor",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Profesor eliminado correctamente",
	})
}
