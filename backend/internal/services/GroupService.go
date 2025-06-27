package services

import (
	"fmt"
	"log"
	"net/http"

	"backend.com/backend/internal/db"
	"backend.com/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func GetAllGroups(c echo.Context) error {
	var group []model.Group
	fmt.Println("DB init:", db.DB)
	if err := db.DB.Find(&group).Error; err != nil {

		log.Printf("Error al obtener grupos: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error interno",
			"error":   err.Error(),
		})

	}

	return c.JSON(http.StatusOK, group)
}
