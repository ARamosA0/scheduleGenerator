package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

func ReadExcelFile(path string) ([]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		log.Fatal("No se encontraron hojas en el archivo")
		return nil, nil
	}
	sheetName := sheets[0]

	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Error al leer filas: %v", err)
		return nil, nil
	}

	if len(rows) == 0 {
		log.Fatal("El archivo está vacío")
		return nil, nil
	}

	columnNames := rows[0]

	return columnNames, nil
}

func ValidateFile(c echo.Context) error {
	// Obtener el archivo del form-data
	fileHeader, err := c.FormFile("file")
	println("FILE", fileHeader)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No se pudo obtener el archivo",
		})
	}

	// Abrir el archivo
	file, err := fileHeader.Open()
	println("FILE2", file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "No se pudo abrir el archivo",
		})
	}
	defer file.Close()

	// Guardar archivo temporalmente
	fileID := uuid.NewString() // github.com/google/uuid
	tempFilePath := "temp/" + fileID + "_" + fileHeader.Filename
	println("FILE3", tempFilePath)
	out, err := os.Create(tempFilePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "No se pudo crear el archivo temporal",
		})
	}
	defer out.Close()
	print("FILE OUT", out)
	// Copiar contenido del archivo a disco
	_, err = io.Copy(out, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "No se pudo copiar el archivo",
		})
	}

	// Leer columnas desde el archivo con tu función
	println("TEMPFILEPATH", tempFilePath)
	columnNames, err := ReadExcelFile(tempFilePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al leer el archivo Excel",
		})
	}

	// Responder con los nombres de columnas
	return c.JSON(http.StatusOK, map[string]interface{}{
		"columns": columnNames,
		"fileId":  tempFilePath,
	})
}

func ExtractData(fileID string) ([]map[string]string, error) {
	println("------ EXTRACT DATA --------", fileID)
	// Buscar archivo en la carpeta temp/
	files, err := filepath.Glob(fileID)
	if err != nil || len(files) == 0 {
		println("Error al buscar el archivo")
	}

	filePath := files[0]

	// Leer el archivo
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %w", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("el archivo no tiene hojas")
	}
	sheetName := sheets[0]

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("error al leer filas: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("el archivo no contiene datos suficientes")
	}

	// Asumimos que la primera fila son los nombres de las columnas
	headers := rows[0]
	dataRows := rows[1:]

	// Convertir las filas en objetos con claves de columna
	var result []map[string]string
	for _, row := range dataRows {
		rowMap := map[string]string{}
		for i, cell := range row {
			if i < len(headers) {
				rowMap[headers[i]] = cell
			}
		}
		result = append(result, rowMap)
	}
	jsonData, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println("FILE ROWS:\n", string(jsonData))
	return result, nil
}
