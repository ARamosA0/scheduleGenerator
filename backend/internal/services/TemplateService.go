package services

import (
	"fmt"
	"net/http"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func GetAllTemplates(c echo.Context) error {
	var templates []model.Template

	if err := db.DB.Find(&templates).Error; err != nil {
		println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al obtener salas",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, templates)
}

func GetTemplate(c echo.Context) error {
	id := c.Param("id")
	var room model.Room

	if err := db.DB.Find(&room, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener template"})
	}

	return c.JSON(http.StatusOK, room)
}

func CreateTemplate(c echo.Context) error {
	var t model.Template

	fmt.Printf("TEMPLATE RECIBIDO %+v\n", t)
	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	fmt.Printf("TEMPLATE RECIBIDO %+v\n", t)
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
	var template model.Template
	if err := db.DB.First(&template, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Template no encontrado"})
	}

	var input model.Template
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	template.Name = input.Name
	template.DaysRange = input.DaysRange

	fmt.Printf("TEMAPLTE UPDATE ---------------------- ", template)
	if err := db.DB.Save(&template).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo actualizar"})
	}

	return c.JSON(http.StatusCreated, template)
}

func DeleteTemplate(c echo.Context) error {
	id := c.Param("id")

	var room model.Room
	if err := db.DB.First(&room, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Template no encontrado",
		})
	}

	if err := db.DB.Delete(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar el template",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Profesor eliminado correctamente",
	})
}
