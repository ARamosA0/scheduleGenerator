package services

import (
	"fmt"
	"net/http"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

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

	if t.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'name' es obligatorio"})
	}

	fmt.Printf("TEACHER RECIBIDO %+v\n", t)
	fmt.Printf("AvailableDays: %+v\n", t.AvailableDays)

	if err := db.DB.Create(&t).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el profesor"})
	}

	return c.JSON(http.StatusCreated, t)
}

func UpdateTeacher(c echo.Context) error {
	id := c.Param("id")
	var teacher model.Teacher
	if err := db.DB.First(&teacher, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Profesor no encontrado"})
	}

	var input model.Teacher
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	teacher.Name = input.Name
	teacher.LastName = input.LastName
	teacher.Email = input.Email
	teacher.Phone = input.Phone
	teacher.Specialty = input.Specialty
	teacher.AvailableDays = input.AvailableDays

	if err := db.DB.Save(&teacher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo actualizar"})
	}

	return c.JSON(http.StatusCreated, teacher)
}

func DeleteTeacher(c echo.Context) error {
	id := c.Param("id")

	var teacher model.Teacher
	if err := db.DB.First(&teacher, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Profesor no encontrado",
		})
	}

	if err := db.DB.Delete(&teacher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar el profesor",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Profesor eliminado correctamente",
	})
}
