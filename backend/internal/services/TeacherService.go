package services

import (
	"fmt"
	"net/http"
	"strings"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func UploadDataFromExcel(c echo.Context) error {

	var body model.MappingBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error al leer el cuerpo de la solicitud",
		})
	}
	println("MAPPING BODY FILE", body.FileID)

	data, err := ExtractData(body.FileID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al leer el archivo Excel",
			"details": err.Error(),
		})
	}
	// println("DATA EXT$RACT", data)

	// Elimina la data
	if err := db.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Teacher{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar datos anteriores",
		})
	}

	// Inserta la data
	for _, teacher := range data {

		var t model.Teacher

		for _, mapping := range body.SelectedMapping {
			value := teacher[mapping.Document]

			switch mapping.Database {
			case "name":
				t.Name = value
			case "lastName":
				t.LastName = value
			case "email":
				t.Email = value
			case "phone":
				t.Phone = value
			case "specialty":
				t.Specialty = value
			case "available_days":
				var availableDays model.AvailableDays

				if strings.Contains(value, ",") {
					parts := strings.Split(value, ",")
					for _, p := range parts {
						day := strings.ToUpper(strings.TrimSpace(p))
						availableDays = append(availableDays, day)
					}
				} else {
					availableDays = append(availableDays, strings.ToUpper(strings.TrimSpace(value)))
				}

				t.AvailableDays = availableDays
			}
		}

		if err := db.DB.Create(&t).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el profesor"})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Datos cargados correctamente",
	})
}
