package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo leer el cuerpo"})
	}
	fmt.Println("CUERPO DEL REQUEST:", string(body))
	c.Request().Body = io.NopCloser(bytes.NewReader(body))

	var process model.Process
	if err := c.Bind(&process); err != nil {
		fmt.Println("ERROR DE BIND:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
	}

	var template model.Template
	if err := db.DB.First(&template, process.SelectedData.SelectedTemplate).Error; err != nil {
		println(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Error al obtener template",
			"details": err.Error(),
		})
	}

	fmt.Printf("TEMPLATE ------------------ %+v\n", template)
	fmt.Printf("PROCESO %+v\n", process)

	subjectsJSON, _ := json.Marshal(process.SelectedData.SelectedSubjects)
	teachersJSON, _ := json.Marshal(process.SelectedData.SelectedTeachers)
	roomsJSON, _ := json.Marshal(process.SelectedData.SelectedRooms)
	templateJSON, _ := json.Marshal(template)

	data := model.Assignment{
		Template:    templateJSON,
		Subjects:    subjectsJSON,
		Teachers:    teachersJSON,
		Rooms:       roomsJSON,
		ProcessName: process.ProcessData.ProcessName,
		Population:  process.ProcessData.Population,
		Generations: process.ProcessData.Generations,
		Mutation:    process.ProcessData.Mutation,
		CrossOver:   process.ProcessData.CrossOver,
		Selection:   process.ProcessData.Selection,
		Reinsertion: process.ProcessData.Reinsertion,
	}

	if err := db.DB.Create(&data).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el proceso"})
	}

	// API DE RUST
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("ERROR SERIALIZANDO JSON PARA RUST:", err)
	} else {
		resp, err := http.Post("http://schedulegenerator-algoritm-service-1:8088/generar", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("ERROR LLAMANDO A API RUST:", err)
		} else {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Println("RESPUESTA DE RUST:", string(body))
		}
	}
	// schedule := model.Schedule{
	// 	Assignment_id: data.ID,
	// 	StartDate: jsonData.,
	// }

	// if err := db.DB.Create().Error; err != nil {

	// }

	return c.JSON(http.StatusCreated, data)
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
