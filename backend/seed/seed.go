package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
)

type SeedData struct {
	Group      []model.Group      `json:"Group"`
	Room       []model.Room       `json:"Room"`
	RoomType   []model.RoomType   `json:"Room"`
	Rule       []model.Rule       `json:"Rule"`
	Subject    []model.Subject    `json:"Subject"`
	Teacher    []model.Teacher    `json:"Teacher"`
	Assignment []model.Assignment `json:"Assignment"`
}

func main() {

	dbConn, err := db.GetConnection()
	if err != nil {
		log.Fatalf("Error conectando a DB: %v", err)
	}

	file, err := os.ReadFile("data.json")
	if err != nil {
		panic(fmt.Sprintf("Error leyendo JSON: %v", err))
	}

	var data SeedData
	if err := json.Unmarshal(file, &data); err != nil {
		panic(fmt.Sprintf("Error parseando JSON: %v", err))
	}

	// Insertar datos
	if len(data.Group) > 0 {
		dbConn.Create(&data.Group)
	}
	if len(data.Room) > 0 {
		dbConn.Create(&data.Room)
	}
	if len(data.RoomType) > 0 {
		dbConn.Create(&data.RoomType)
	}
	if len(data.Rule) > 0 {
		dbConn.Create(&data.Rule)
	}
	if len(data.Subject) > 0 {
		dbConn.Create(&data.Subject)
	}
	if len(data.Teacher) > 0 {
		dbConn.Create(&data.Teacher)
	}
	if len(data.Assignment) > 0 {
		dbConn.Create(&data.Assignment)
	}

	fmt.Println("Seed data insertado correctamente.")
}
