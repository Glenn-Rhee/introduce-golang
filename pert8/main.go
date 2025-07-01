package main

import (
	"fmt"
	"log"
	"net/http"

	"pert8_10123178/configs"
	"pert8_10123178/handlers"
	"pert8_10123178/middlewares"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// TODO : ACT NO 2
	PORT := 28081 // Contoh: 8081, ganti dengan 4 digit terakhir NPM Anda

	// Inisialisasi koneksi database
	configs.ConnectDB()
	if configs.DB == nil {
		log.Fatal("Koneksi database gagal")
	}
	defer func() {
		if err := configs.DB.Close(); err != nil {
			log.Printf("Error menutup koneksi database: %v", err)
		}
	}()

	// Inisialisasi server dan route
	mux := http.NewServeMux()

	// File server untuk konten statis
	fileServer := http.FileServer(http.Dir("catalog"))

	// Route untuk file statis
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeStaticFile(w, r, "catalog", fileServer)
	})

	// TODO : ACT NO 2
	mux.HandleFunc("/api/events/", handlers.HandleEvents)
	
	// Terapkan middleware logging
	loggedMux := middlewares.LogRequestHandler(mux)

	fmt.Printf("Server berjalan di http://localhost:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), loggedMux))
}
