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

func GetAllRoom(c echo.Context) error {
	var room []model.Room

	if err := db.DB.Find(&room).Error; err != nil {
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

	fmt.Printf("SALON RECIBIDO %+v\n", r)

	// if r.Code == "" {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'name' es obligatorio"})
	// }

	if err := db.DB.Create(&r).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el curso"})
	}

	return c.JSON(http.StatusCreated, r)
}

func UpdateRoom(c echo.Context) error {
	id := c.Param("id")
	var room model.Room
	if err := db.DB.First(&room, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Profesor no encontrado"})
	}

	fmt.Printf("UPDATE SALON RECIBIDO %+v\n", room)

	var input model.Room
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	room.Code = input.Code
	room.Name = input.Name
	room.Capacity = input.Capacity
	room.RoomType = input.RoomType
	room.Floor = input.Floor
	room.Building = input.Building
	room.Observations = input.Observations

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

func UploadRoomFromExcel(c echo.Context) error {

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
	if err := db.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Room{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al eliminar datos anteriores",
		})
	}

	// Inserta la data
	for _, teacher := range data {

		var s model.Room

		for _, mapping := range body.SelectedMapping {
			value := teacher[mapping.Document]

			switch mapping.Database {
			case "code":
				s.Code = value
			case "name":
				s.Name = value
			case "capacity":
				capacity, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("Error al convertir '%s' a int para 'semester': %v", value, err)
					capacity = 0
				}
				s.Capacity = capacity
			case "room_type":
				room_type, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("Error al convertir '%s' a int para 'semester': %v", value, err)
					room_type = 0
				}
				s.RoomType = room_type
			case "floor":
				floor, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("Error al convertir '%s' a int para 'semester': %v", value, err)
					floor = 0
				}
				s.Floor = floor
			case "building":
				s.Building = value
			case "observations":
				s.Observations = value
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
