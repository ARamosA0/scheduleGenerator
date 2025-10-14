package services

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

var (
	clients   = make(map[uint]chan string) // uint = assignment_id
	clientsMu sync.Mutex                   // mutex para acceso concurrente
)

func HandleServerEvents(c echo.Context) error {
	assignmentIDParam := c.QueryParam("id")
	var assignmentID uint
	fmt.Sscanf(assignmentIDParam, "%d", &assignmentID)

	fmt.Printf(" Cliente conectado para assignment ID: %d\n", assignmentID)

	messageChan := make(chan string, 10) // Buffer importante

	clientsMu.Lock()
	clients[assignmentID] = messageChan
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(clients, assignmentID)
		clientsMu.Unlock()
		close(messageChan)
		fmt.Printf(" Cliente desconectado: %d\n", assignmentID)
	}()

	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().WriteHeader(http.StatusOK)

	// Mensaje inicial para confirmar conexión
	fmt.Fprintf(c.Response(), "data: Conectado exitosamente\n\n")
	c.Response().Flush()

	for {
		select {
		case <-c.Request().Context().Done():
			fmt.Println("️Contexto cancelado")
			return nil
		case msg, ok := <-messageChan:
			if !ok {
				fmt.Println("Canal cerrado")
				return nil
			}
			fmt.Printf("Enviando mensaje: %s\n", msg)
			fmt.Fprintf(c.Response(), "data: %s\n\n", msg)
			c.Response().Flush()

			if msg == "__CLOSE__" {
				fmt.Println("Cerrando conexión por __CLOSE__")
				return nil
			}
		}
	}
}
