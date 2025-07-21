package services

import (
	"fmt"
	"net/http"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func GetAllAssigments(c echo.Context) error {
	var assigment []model.Assignment

	if err := db.DB.Find(&assigment).Error; err != nil {
		println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al obtener salas",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, assigment)
}

func GetAssigment(c echo.Context) error {
	var assigment model.Assignment

	if err := db.DB.Find(&assigment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener profesores"})
	}

	return c.JSON(http.StatusOK, assigment)
}

func CreateAssigment(c echo.Context) error {
	var a model.Assignment
	if err := c.Bind(&a); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	fmt.Printf("RECIBIDO %+v\n", a)

	if err := db.DB.Create(&a).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el curso"})
	}

	return c.JSON(http.StatusCreated, a)
}

func UpdateAssigment(c echo.Context) error {
	id := c.Param("id")
	var assigment model.Assignment
	if err := db.DB.First(&assigment, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Profesor no encontrado"})
	}

	fmt.Printf("UPDATE SALON RECIBIDO %+v\n", assigment)

	var input model.Room
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	// assigment.Code = input.Code
	// room.Name = input.Name
	// room.Capacity = input.Capacity
	// room.RoomType = input.RoomType
	// room.Floor = input.Floor
	// room.Building = input.Building
	// room.Observations = input.Observations

	if err := db.DB.Save(&assigment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo actualizar"})
	}

	return c.JSON(http.StatusCreated, assigment)
}

func DeleteAssigment(c echo.Context) error {
	id := c.Param("id")

	var assigment model.Assignment
	if err := db.DB.First(&assigment, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Profesor no encontrado",
		})
	}

	if err := db.DB.Delete(&assigment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar el profesor",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Profesor eliminado correctamente",
	})
}
