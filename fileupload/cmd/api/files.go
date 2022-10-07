/*
	CMPS4191 - Quiz #2
	Rene Sanchez - 2018118383
*/
// Filename: cmd/api/files.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"fileupload.renesanchez455.net/internal/data"
	"fileupload.renesanchez455.net/internal/validator"
)

// createEntryHandler for the "POST /v1/user" endpoint
func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name       string `json:"name"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		Address    string `json:"address"`
		Occupation string `json:"occupation"`
	}
	// Initialize a new json.Decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Initialize a new Validator Instance
	v := validator.New()
	// Use the Check() method to execute our validation checks
	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(input.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(input.Phone, validator.PhoneRx), "phone", "must be a valid phone number")

	v.Check(input.Email != "", "email", "must be provided")
	v.Check(validator.Matches(input.Email, validator.EmailRx), "email", "must be a valid email address")

	v.Check(input.Address != "", "address", "must be provided")
	v.Check(len(input.Address) <= 500, "address", "must not be more than 500 bytes long")

	v.Check(input.Occupation != "", "occupation", "must be provided")
	v.Check(len(input.Occupation) <= 200, "occupation", "must not be more than 200 bytes long")

	// Check the map to determine if there were any validation errors
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Display the request
	fmt.Fprintf(w, "%+v\n", input)
}

// showEntryHandler for the "GET /v1/rand/:id" endpoint
func (app *application) showRandStringHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Display the random string
	num := int(id) //Convert int64 id to int
	fmt.Fprintf(w, "show random string for %d\n", num)
	fmt.Fprintln(w, app.Tools.generateRandomString(num))
}

// showinputHandler for the "GET /v1/user/:id" endpoint
func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the input struct containing the ID we extracted
	// from our URL and some sample data
	user := data.User{
		ID:         id,
		CreatedAt:  time.Now(),
		Name:       "Rene Sanchez",
		Phone:      "623-5411",
		Address:    "20 Caterpillar street",
		Occupation: "Student",
		Version:    1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
