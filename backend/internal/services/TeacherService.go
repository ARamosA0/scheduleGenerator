package services

import (
	"net/http"
	"time"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

var teachers = []model.Teacher{
	{
		Name:           "Teacher1",
		AvailableTimes: []time.Time{time.Now()}, // Use slice initialization
	},
	{
		Name:           "Teacher2",
		AvailableTimes: []time.Time{time.Now()}, // Use slice initialization
	},
}

var teacher = model.Teacher{
	Name:           "Teacher1",
	AvailableTimes: []time.Time{time.Now()},
}

func GetAllTeachers(c echo.Context) error {
	var teachers []model.Teacher

	if err := db.DB.Find(&teachers).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al obtener profesores",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, teachers)
}

func GetTeacher(c echo.Context) error {
	var teacher model.Teacher

	if err := db.DB.Find(&teacher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener profesores"})
	}

	return c.JSON(http.StatusOK, teacher)
}

func CreateTeacher(c echo.Context) error {
	var t model.Teacher
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	// Validación opcional
	if t.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'name' es obligatorio"})
	}

	// Guardar en la base de datos
	if err := db.DB.Create(&t).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el profesor"})
	}

	// Devolver el objeto creado
	return c.JSON(http.StatusCreated, t)
}
