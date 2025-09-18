package services

import (
	"log"
	"net/http"
	"strconv"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	subject.Specialty = input.Specialty

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

func UploadSubjectFromExcel(c echo.Context) error {

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

	// Elimina la data
	if err := db.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Subject{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar datos anteriores",
		})
	}

	// Inserta la data
	for _, teacher := range data {

		var s model.Subject

		for _, mapping := range body.SelectedMapping {
			value := teacher[mapping.Document]

			switch mapping.Database {
			case "code":
				s.Code = value
			case "name":
				s.Name = value
			case "credits":
				credits, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("Error al convertir '%s' a int para 'credits': %v", value, err)
					credits = 0
				}
				s.Credits = credits
			case "hours":
				hours, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("Error al convertir '%s' a int para 'hours': %v", value, err)
					hours = 0
				}
				s.Hours = hours
			case "semester":
				semester, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("Error al convertir '%s' a int para 'semester': %v", value, err)
					semester = 0
				}
				s.Semester = semester
			case "career":
				s.Career = value
			case "requirements":
				s.Requirements = value
			case "description":
				s.Description = value
			case "required_room_type":
				roomType, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("Error al convertir '%s' a int para 'semester': %v", value, err)
					roomType = 0
				}
				s.RequiredRoomType = roomType
			case "specialty":
				// t.Specialty = value
				specialty, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("Error al convertir '%s' a int para 'credits': %v", value, err)
					specialty = 0
				}
				s.Specialty = specialty
			}

		}

		if err := db.DB.Create(&s).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el profesor"})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Datos cargados correctamente",
	})
}
