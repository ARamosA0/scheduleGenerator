package services

import (
	"net/http"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func GetAllSubject(c echo.Context) error {
	var subject []model.Subject

	if err := db.DB.Find(&subject).Error; err != nil {
		println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al obtener salas",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, subject)
}

func GetSubject(c echo.Context) error {
	var subject model.Subject

	if err := db.DB.Find(&subject).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener profesores"})
	}

	return c.JSON(http.StatusOK, subject)
}

func CreateSubject(c echo.Context) error {
	var s model.Subject
	if err := c.Bind(&s); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	if s.Code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'name' es obligatorio"})
	}

	if err := db.DB.Create(&s).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el curso"})
	}

	return c.JSON(http.StatusCreated, s)
}

func UpdateSubject(c echo.Context) error {
	id := c.Param("id")
	var subject model.Subject
	if err := db.DB.First(&subject, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Profesor no encontrado"})
	}

	var input model.Subject
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	subject.Code = input.Code
	subject.Name = input.Name
	subject.Credits = input.Credits
	subject.Hours = input.Hours
	subject.Semester = input.Semester
	subject.Career = input.Career
	subject.Requirements = input.Requirements
	subject.RequiredRoomType = input.RequiredRoomType

	if err := db.DB.Save(&subject).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo actualizar"})
	}

	return c.JSON(http.StatusCreated, subject)
}

func DeleteSubjject(c echo.Context) error {
	id := c.Param("id")

	var subject model.Subject
	if err := db.DB.First(&subject, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Profesor no encontrado",
		})
	}

	if err := db.DB.Delete(&subject).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar el profesor",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Profesor eliminado correctamente",
	})
}
