package model

type ColumnMapping struct {
	Document string `json:"document"`
	Database string `json:"database"`
}

type MappingBody struct {
	FileID          string          `json:"fileId"`
	SelectedMapping []ColumnMapping `json:"selectedMapping"`
}
