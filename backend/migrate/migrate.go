package main

import (
	"log"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
)

func main() {
	log.Println("Ejecutando las migraciones ....")
	dbConn, err := db.GetConnection()
	if err != nil {
		log.Fatalf("Error conectando a DB: %v", err)
	}

	err = dbConn.AutoMigrate(
		&model.Assignment{},
		&model.Group{},
		&model.Room{},
		&model.RoomType{},
		&model.Rule{},
		&model.Subject{},
		&model.Teacher{},
	)

	if err != nil {
		log.Fatalf("Error en migraciones: %v", err)
	}
	log.Println("MIgraciones completadas!")
}
