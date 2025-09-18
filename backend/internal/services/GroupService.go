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
	var groups []model.Group

	if err := db.DB.Preload("Subjects").Find(&groups).Error; err != nil {
		println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al obtener salas",
			"details": err.Error(),
		})
	}

	var response []model.GroupInput

	for _, group := range groups {
		subjectIDs := make([]uint, len(group.Subjects))
		fmt.Println("SUBJECTIDS", subjectIDs)
		for i, subject := range group.Subjects {
			subjectIDs[i] = subject.ID
		}

		response = append(response, model.GroupInput{
			ID:         group.ID,
			Name:       group.Name,
			Size:       group.Size,
			SubjectIDs: subjectIDs,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func GetGroup(c echo.Context) error {
	var group model.Group

	if err := db.DB.Find(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener profesores"})
	}

	return c.JSON(http.StatusOK, group)
}

func CreateGroup(c echo.Context) error {
	var gi model.GroupInput
	if err := c.Bind(&gi); err != nil {
		fmt.Println("GROUP ERROR", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	group := model.Group{
		Name: gi.Name,
		Size: gi.Size,
	}

	var subjects []model.Subject

	fmt.Println("GROUP CREATE", group)

	if len(gi.SubjectIDs) > 0 {
		if err := db.DB.Find(&subjects, gi.SubjectIDs).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al buscar las materias"})
		}
	}

	group.Subjects = subjects

	if err := db.DB.Create(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el curso"})
	}

	return c.JSON(http.StatusCreated, group)
}

func UpdateGroup(c echo.Context) error {
	id := c.Param("id")
	var group model.Group
	if err := db.DB.First(&group, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "GRupo no encontrado"})
	}

	var input model.Group
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	group.Name = input.Name
	group.Size = input.Size
	// group.Subjects = input.Subjects

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
