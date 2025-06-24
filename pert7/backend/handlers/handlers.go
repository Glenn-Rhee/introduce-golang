package handlers

import (
	"encoding/json"
	"net/http"
	"pert7/backend/models"
	"strconv"
)

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetNotes(w, r)
	} else if r.Method == http.MethodPost {
		AddNote(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		DeleteNote(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Notes)
}

func AddNote(w http.ResponseWriter, r *http.Request) {
	var newNote models.Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newNote.ID = len(models.Notes) + 1

	models.Notes = append(models.Notes, newNote)
	models.SaveNotes()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNote)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, note := range models.Notes {
		if note.ID == id {
			models.Notes = append(models.Notes[:i], models.Notes[i+1:]...)
			models.SaveNotes()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Note deleted"))
			return
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}
