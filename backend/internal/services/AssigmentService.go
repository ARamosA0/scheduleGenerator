package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
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
	c.Request().Body = io.NopCloser(bytes.NewReader(body))

	// fmt.Println("RAW BODY:", string(body))

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

	subjectsJSON, _ := json.Marshal(process.SelectedData.SelectedSubjects)
	teachersJSON, _ := json.Marshal(process.SelectedData.SelectedTeachers)
	roomsJSON, _ := json.Marshal(process.SelectedData.SelectedRooms)
	groupsJSON, _ := json.Marshal(process.SelectedData.SelectedGroups)
	templateJSON, _ := json.Marshal(template)

	// fmt.Println("CONFIG", process.ProcessData.Generations)

	data := model.Assignment{
		Template:    templateJSON,
		Subjects:    subjectsJSON,
		Teachers:    teachersJSON,
		Rooms:       roomsJSON,
		Groups:      groupsJSON,
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

	// Enviar la tarea al algoritmo de forma asíncrona
	go func(assignment model.Assignment) {
		fmt.Printf(" Iniciando proceso para assignment %d\n", assignment.ID)

		sendEvent := func(id uint, msg string, status string, progress uint, scheduleId string) {
			// Reintentar hasta 10 veces (5 segundos total)
			for i := 0; i < 10; i++ {
				clientsMu.Lock()
				ch, ok := clients[1]
				clientsMu.Unlock()

				if ok {
					fmt.Printf(" Enviando a cliente %d: %s\n", 1, msg)
					event := model.ProcessEventMessage{
						Message:  msg,
						Progress: progress,
						Status:   status,
					}
					data, _ := json.Marshal(event)
					ch <- string(data)
					return
				}

				if i == 0 {
					fmt.Printf(" Esperando cliente %d... intento %d/10\n", 1, i+1)
				}
				time.Sleep(500 * time.Millisecond)
			}

			fmt.Printf("️ No hay cliente conectado para ID %d después de esperar\n", id)
		}

		sendEvent(assignment.ID, "Proceso iniciado ", "__START__", 0, "")

		jsonData, _ := json.Marshal(assignment)
		resp, err := http.Post("http://algoritm-service:8088/generar", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			sendEvent(assignment.ID, fmt.Sprintf(" Error enviando a Rust: %v", err), "__ERROR__", 100, "")
			db.DB.Model(&assignment).Update("Status", "error")
			return
		}
		defer resp.Body.Close()

		sendEvent(assignment.ID, "Procesando algoritmo en Rust...", "__START__", 0, "")

		body, _ := io.ReadAll(resp.Body)
		var rustResult model.ScheduleResponse
		if err := json.Unmarshal(body, &rustResult); err != nil {
			sendEvent(assignment.ID, fmt.Sprintf(" Error decodificando respuesta: %v", err), "__ERROR__", 100, "")
			db.DB.Model(&assignment).Update("Status", "error")
			return
		}

		jsonBytes, _ := json.Marshal(rustResult)
		schedule := model.Schedule{
			Assignment_id:    uint(assignment.ID),
			ScheduleResponse: datatypes.JSON(jsonBytes),
		}
		db.DB.Create(&schedule)

		scheduleIDStr := fmt.Sprintf("%d", schedule.ID)
		sendEvent(assignment.ID, "Proceso completado correctamente.", "__END__", 100, scheduleIDStr)
		sendEvent(assignment.ID, "__CLOSE__", "__END__", 100, scheduleIDStr)
	}(data)

	// Responder inmediatamente al frontend
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message":    "Proceso iniciado",
		"process_id": data.ID,
		"status":     "pending",
	})

	// API DE RUST
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println("ERROR SERIALIZANDO JSON PARA RUST:", err)
	// 	return err
	// }
	// // Enviar a Rust
	// resp, err := http.Post("http://schedulegenerator-algoritm-service-1:8088/generar", "application/json", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	fmt.Println("ERROR LLAMANDO A API RUST:", err)
	// 	return err
	// }
	// defer resp.Body.Close()

	// body, bodyErr := io.ReadAll(resp.Body)
	// if bodyErr != nil {
	// 	fmt.Println("ERROR LEYENDO RESPUESTA:", bodyErr)
	// 	return bodyErr
	// }

	// var rustResult model.ScheduleResponse
	// err = json.Unmarshal(body, &rustResult)
	// if err != nil {
	// 	fmt.Println("ERROR DECODIFICANDO JSON DE RUST:", err)
	// 	return err
	// }

	// fmt.Println("RUST SCHEDULE", rustResult)

	// fmt.Println("--------------------------------------------")
	// fmt.Println("RUST RESULT :", rustResult)
	// fmt.Println("Schedule RESPONSE :", rustResult)
	// fmt.Println("--------------------------------------------")

	// jsonBytes, err := json.Marshal(rustResult)
	// if err != nil {
	// 	return err
	// }

	// schedule := model.Schedule{
	// 	Assignment_id:    uint(data.ID),
	// 	ScheduleResponse: datatypes.JSON(jsonBytes),
	// }

	// if err := db.DB.Create(&schedule).Error; err != nil {
	// 	fmt.Println("Error insertando en DB:", err)
	// } else {
	// 	fmt.Println("Schedule insertado:", schedule)
	// }

	// rustResult.ScheduleId = schedule.ID

	// return c.JSON(http.StatusCreated, rustResult)
}

func UpdateAssigment(c echo.Context) error {
	id := c.Param("id")
	var assigment model.Assignment
	if err := db.DB.First(&assigment, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Profesor no encontrado"})
	}

	// fmt.Printf("UPDATE SALON RECIBIDO %+v\n", assigment)

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
