/*
	CMPS4191 - Quiz #2
	Rene Sanchez - 2018118383
*/
// Filename: cmd/api/files.go

package main

import (
	"fmt"
	"net/http"

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

	// Copy the values from the input struct to a new User struct
	user := &data.User{
		Name:       input.Name,
		Phone:      input.Phone,
		Email:      input.Email,
		Address:    input.Address,
		Occupation: input.Occupation,
	}
	// Initialize a new Validator Instance
	v := validator.New()

	// Check the map to determine if there were any validation errors
	if data.ValidateUser(v, user); !v.Valid() {
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
	fmt.Fprintln(w, app.Tools.generateRandomString(num))
}
