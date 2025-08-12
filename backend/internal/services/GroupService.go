package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllGroups(c echo.Context) error {
	var group []model.Group

	if err := db.DB.Find(&group).Error; err != nil {
		println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al obtener salas",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, group)
}

func GetGroup(c echo.Context) error {
	var group model.Group

	if err := db.DB.Find(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener profesores"})
	}

	return c.JSON(http.StatusOK, group)
}

func CreateGroup(c echo.Context) error {
	var g model.Group
	if err := c.Bind(&g); err != nil {
		fmt.Println("GROUP ERROR", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	fmt.Println("GROUP CREATE", g)

	if g.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'name' es obligatorio"})
	}

	if err := db.DB.Create(&g).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el curso"})
	}

	return c.JSON(http.StatusCreated, g)
}

func UpdateGroup(c echo.Context) error {
	id := c.Param("id")
	var group model.Group
	if err := db.DB.First(&group, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Profesor no encontrado"})
	}

	var input model.Group
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	group.Name = input.Name
	group.Size = input.Size

	if err := db.DB.Save(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo actualizar"})
	}

	return c.JSON(http.StatusCreated, group)
}

func DeleteGroup(c echo.Context) error {
	id := c.Param("id")

	var group model.Group
	if err := db.DB.First(&group, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Profesor no encontrado",
		})
	}

	if err := db.DB.Delete(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar el profesor",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Profesor eliminado correctamente",
	})
}

func UploadGroupFromExcel(c echo.Context) error {

	var body model.MappingBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error al leer el cuerpo de la solicitud",
		})
	}

	data, err := ExtractData(body.FileID)

	if err != nil {
		print("ERROR LUEGO DE EXTRACT DATA", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al leer el archivo Excel",
			"details": err.Error(),
		})
	}

	// Elimina la data
	if err := db.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Group{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar datos anteriores",
		})
	}

	// Inserta la data
	for _, group := range data {

		var g model.Group

		for _, mapping := range body.SelectedMapping {
			value := group[mapping.Document]

			switch mapping.Database {

			case "name":
				g.Name = value
			case "size":
				size, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("Error al convertir '%s' a int para 'credits': %v", value, err)
					size = 0
				}
				g.Size = size
			}
		}

		if err := db.DB.Create(&g).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el profesor"})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Datos cargados correctamente",
	})
}
