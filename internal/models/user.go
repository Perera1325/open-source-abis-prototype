package models

type User struct {
	ID        string    `json:"id"`
	Embedding []float64 `json:"embedding"`
}
