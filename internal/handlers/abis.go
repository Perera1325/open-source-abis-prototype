package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Perera1325/open-source-abis-prototype/internal/models"
	"github.com/Perera1325/open-source-abis-prototype/internal/storage"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ABIS running"))
}

func Enroll(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	storage.Users[user.ID] = user
	json.NewEncoder(w).Encode(user)
}
