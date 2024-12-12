package controller

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is healthy and ok ")
}
