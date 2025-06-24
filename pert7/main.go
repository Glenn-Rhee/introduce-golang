package main

import (
	"fmt"
	"log"
	"net/http"
	"pert7/backend/handlers"
	"pert7/backend/models"
)

const PORT = 20029

func main() {
	models.LoadNotes()

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/notes", handlers.NotesHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)

	log.Printf("Server running on http:localhost:%d/...", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatalf("Error starting the server on port %d: %v", PORT, err)
	}
}
