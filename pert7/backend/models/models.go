package models

import (
	"encoding/json"
	"os"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Notes []Note

const filePath = "notes.json"

func LoadNotes() {
	file, err := os.ReadFile(filePath)
	if err == nil {
		json.Unmarshal(file, &Notes)
	} else {
		Notes = []Note{}
	}
}

func SaveNotes() {
	file, _ := json.MarshalIndent(Notes, "", "  ")
	os.WriteFile(filePath, file, 0644)
}
