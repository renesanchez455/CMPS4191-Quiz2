// Filename: internal/data/files.go

package data

import (
	"time"
)

type User struct {
	ID         int64
	CreatedAt  time.Time
	Name       string
	Phone      string
	Email      string
	Address    string
	Occupation string
	Version    int32
}
