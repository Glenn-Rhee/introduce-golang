package middlewares

import (
	"log"
	"net/http"
	"time"
)

// LogRequestHandler mencatat setiap request yang masuk ke server
func LogRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
		
		next.ServeHTTP(w, r)
		
		duration := time.Since(start)
		log.Printf("Request completed: %s %s - Duration: %v", r.Method, r.URL.Path, duration)
	})
}