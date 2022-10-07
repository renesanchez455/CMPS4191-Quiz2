// Filename: internal/data/files.go

package data

import (
	"time"
)

type User struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email,omitempty"`
	Address    string    `json:"address"`
	Occupation string    `json:"occupation"`
	Version    int32     `json:"version"`
}
