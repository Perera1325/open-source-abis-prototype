package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/Perera1325/open-source-abis-prototype/internal/models"
	"github.com/Perera1325/open-source-abis-prototype/internal/storage"
)

const similarityThreshold = 0.80

// Health check endpoint
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ABIS running"))
}

// Enroll a new user with embedding
func Enroll(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID == "" || len(user.Embedding) == 0 {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	storage.Mu.Lock()
	storage.Users[user.ID] = user
	storage.Mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Match input embedding against all stored users
func Match(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var input models.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(input.Embedding) == 0 {
		http.Error(w, "Embedding required", http.StatusBadRequest)
		return
	}

	bestScore := -1.0
	var bestMatch models.User

	storage.Mu.RLock()
	userCount := len(storage.Users)

	for _, user := range storage.Users {
		score := cosineSimilarity(input.Embedding, user.Embedding)

		if score > bestScore {
			bestScore = score
			bestMatch = user
		}
	}
	storage.Mu.RUnlock()

	duration := time.Since(start)

	if bestScore < similarityThreshold {
		http.Error(w, "No strong match found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"matched_user": bestMatch,
		"similarity":   bestScore,
		"threshold":    similarityThreshold,
		"users_scanned": userCount,
		"search_time":   duration.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Cosine similarity calculation
func cosineSimilarity(a, b []float64) float64 {
	if len(a) != len(b) {
		return -1
	}

	var dotProduct float64
	var normA float64
	var normB float64

	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return -1
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}
