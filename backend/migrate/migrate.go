package main

import (
	"fmt"
	"log"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"gorm.io/gorm"
)

func ResetDatabase(db *gorm.DB) error {
	models := []interface{}{
		&model.Assignment{},
		&model.Group{},
		&model.Room{},
		&model.RoomType{},
		&model.Rule{},
		&model.Subject{},
		&model.Teacher{},
	}

	for _, m := range models {
		if db.Migrator().HasTable(m) {
			if err := db.Migrator().DropTable(m); err != nil {
				return fmt.Errorf("error al eliminar tabla: %w", err)
			}
		}
	}

	return db.AutoMigrate(models...)
}

func main() {
	log.Println("Ejecutando las migraciones ....")
	dbConn, err := db.GetConnection()
	if err != nil {
		log.Fatalf("Error conectando a DB: %v", err)
	}

	ResetDatabase(dbConn)

	if err != nil {
		log.Fatalf("Error en migraciones: %v", err)
	}
	log.Println("MIgraciones completadas!")
}
