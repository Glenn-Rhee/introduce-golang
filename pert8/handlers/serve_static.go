package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// ServeStaticFile melayani file statis dari folder catalog
func ServeStaticFile(w http.ResponseWriter, r *http.Request, baseDir string, fileServer http.Handler) {
	// Hapus slash di awal jika ada
	path := strings.TrimPrefix(r.URL.Path, "/")
	fullPath := filepath.Join(baseDir, path)

	// Jangan layani permintaan ke API melalui file server
	if strings.HasPrefix(r.URL.Path, "/api/") {
		http.NotFound(w, r)
		return
	}

	// Periksa apakah file ada
	_, err := os.Stat(fullPath)
	if err != nil {
		// Jika file tidak ditemukan, alihkan ke halaman utama
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	
	// Layani file statis
	fileServer.ServeHTTP(w, r)
}