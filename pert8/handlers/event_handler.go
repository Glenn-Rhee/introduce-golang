package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"pert8_10123178/configs"
	"pert8_10123178/models"
	"pert8_10123178/utils"
)

// HandleEvents mengarahkan request ke handler yang sesuai berdasarkan method HTTP
func HandleEvents(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var id int
	var err error

	// Ekstrak ID dari URL jika ada
	if len(path) > len("/api/events/") {
		idStr := path[len("/api/events/"):]
		id, err = strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID tidak valid", http.StatusBadRequest)
			return
		}
	} else {
		id = 0
	}

	// Mengarahkan ke handler yang sesuai berdasarkan method HTTP
	switch r.Method {
	case http.MethodGet:
		handleGetEvents(w, id)
	case http.MethodPost:
		handlePostEvent(w, r)
	case http.MethodPut:
		handleUpdateEvent(w, r, id)
	case http.MethodDelete:
		handleDeleteEvent(w, id)
	default:
		http.Error(w, "Method tidak didukung", http.StatusMethodNotAllowed)
	}
}

// handleGetEvents mengambil data event (semua event atau event spesifik)
func handleGetEvents(w http.ResponseWriter, id int) {
	if id == 0 {
		// Ambil semua event
		rows, err := configs.DB.Query(`
			SELECT id_event, title, organizer, description, date, location, price, capacity, 
			image_url, status, created_at, updated_at 
			FROM events
		`)
		if err != nil {
			log.Printf("Database error: %v", err)
			http.Error(w, "Gagal mengambil data event", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var events []models.Event
		for rows.Next() {
			var event models.Event
			err := rows.Scan(
				&event.ID, &event.Title, &event.Organizer, &event.Description,
				&event.Date, &event.Location, &event.Price, &event.Capacity,
				&event.ImageURL, &event.Status, &event.CreatedAt, &event.UpdatedAt,
			)
			if err != nil {
				log.Printf("Database error: %v", err)
				http.Error(w, "Gagal memproses data event", http.StatusInternalServerError)
				return
			}
			events = append(events, event)
		}
		utils.RespondJSON(w, events)
	} else {
		// Ambil event spesifik berdasarkan ID
		var event models.Event
		err := configs.DB.QueryRow(`
			SELECT id_event, title, organizer, description, date, location, price, capacity, 
			image_url, status, created_at, updated_at 
			FROM events WHERE id_event = ?
		`, id).Scan(
			&event.ID, &event.Title, &event.Organizer, &event.Description,
			&event.Date, &event.Location, &event.Price, &event.Capacity,
			&event.ImageURL, &event.Status, &event.CreatedAt, &event.UpdatedAt,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Event tidak ditemukan", http.StatusNotFound)
			} else {
				http.Error(w, "Gagal mengambil data event", http.StatusInternalServerError)
			}
			return
		}
		utils.RespondJSON(w, event)
	}
}

// handlePostEvent menambahkan event baru
func handlePostEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	// Validasi data event
	if strings.TrimSpace(event.Title) == "" || strings.TrimSpace(event.Organizer) == "" ||
		strings.TrimSpace(event.Location) == "" || event.Price < 0 || event.Capacity <= 0 {
		http.Error(w, "Data event tidak lengkap atau tidak valid", http.StatusBadRequest)
		return
	}

	// Set nilai default untuk field yang tidak disetel
	if event.Status == "" {
		event.Status = "upcoming"
	}
	now := time.Now()
	event.CreatedAt = now
	event.UpdatedAt = now

	// Simpan event baru
	result, err := configs.DB.Exec(`
		INSERT INTO events (title, organizer, description, date, location, price, capacity, image_url, status) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, event.Title, event.Organizer, event.Description, event.Date, event.Location,
		event.Price, event.Capacity, event.ImageURL, event.Status)

	if err != nil {
		http.Error(w, "Gagal menambahkan event baru", http.StatusInternalServerError)
		return
	}

	// Dapatkan ID yang baru dibuat
	id, _ := result.LastInsertId()
	utils.RespondJSON(w, map[string]interface{}{
		"message": "Event berhasil ditambahkan!",
		"id":      id,
	})
}

// handleUpdateEvent memperbarui data event
func handleUpdateEvent(w http.ResponseWriter, r *http.Request, id int) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	// Validasi data event
	if strings.TrimSpace(event.Title) == "" || strings.TrimSpace(event.Organizer) == "" ||
		strings.TrimSpace(event.Location) == "" || event.Price < 0 || event.Capacity <= 0 {
		http.Error(w, "Data event tidak lengkap atau tidak valid", http.StatusBadRequest)
		return
	}

	// Update data event
	result, err := configs.DB.Exec(`
		UPDATE events 
		SET title=?, organizer=?, description=?, date=?, location=?, price=?, capacity=?, image_url=?, status=?
		WHERE id_event=?
	`, event.Title, event.Organizer, event.Description, event.Date, event.Location,
		event.Price, event.Capacity, event.ImageURL, event.Status, id)

	if err != nil {
		http.Error(w, "Gagal memperbarui data event", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Event dengan ID tersebut tidak ditemukan", http.StatusNotFound)
		return
	}

	utils.RespondJSON(w, map[string]string{
		"message": "Event berhasil diperbarui!",
	})
}

// handleDeleteEvent menghapus event
func handleDeleteEvent(w http.ResponseWriter, id int) {
	result, err := configs.DB.Exec("DELETE FROM events WHERE id_event=?", id)
	if err != nil {
		http.Error(w, "Gagal menghapus event", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Event dengan ID tersebut tidak ditemukan", http.StatusNotFound)
		return
	}

	utils.RespondJSON(w, map[string]string{
		"message": "Event berhasil dihapus!",
	})
}
