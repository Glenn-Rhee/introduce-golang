package utils

import (
	"encoding/json"
	"net/http"
)

// TODO : ACT NO 5 
func RespondJSON(w http.ResponseWriter, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}