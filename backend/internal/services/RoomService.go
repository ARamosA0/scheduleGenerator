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

func GetRoom(c echo.Context) error {
	var room model.Room

	if err := db.DB.Find(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener profesores"})
	}

	return c.JSON(http.StatusOK, room)
}

func CreateRoom(c echo.Context) error {
	var r model.Room
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	// if r.Code == "" {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'name' es obligatorio"})
	// }

	if err := db.DB.Create(&r).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el curso"})
	}

	return r.JSON(http.StatusCreated, r)
}

func UpdateRoom(c echo.Context) error {
	id := c.Param("id")
	var room model.Room
	if err := db.DB.First(&room, id).Error; err != nil {
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
	subject.RequiredRoomTypeID = input.RequiredRoomTypeID

	if err := db.DB.Save(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo actualizar"})
	}

	return c.JSON(http.StatusCreated, room)
}

func DeleteRoom(c echo.Context) error {
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
