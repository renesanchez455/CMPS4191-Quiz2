// Filename: internal/data/files.go

package data

import (
	"time"

	"fileupload.renesanchez455.net/internal/validator"
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

func ValidateUser(v *validator.Validator, user *User) {
	// Use the Check() method to execute our validation checks
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(user.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(user.Phone, validator.PhoneRx), "phone", "must be a valid phone number")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(validator.Matches(user.Email, validator.EmailRx), "email", "must be a valid email address")

	v.Check(user.Address != "", "address", "must be provided")
	v.Check(len(user.Address) <= 500, "address", "must not be more than 500 bytes long")

	v.Check(user.Occupation != "", "occupation", "must be provided")
	v.Check(len(user.Occupation) <= 200, "occupation", "must not be more than 200 bytes long")

}
