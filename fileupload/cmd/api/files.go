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
)

// createEntryHandler for the "POST /v1/files" endpoint
func (app *application) uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "upload a new file..")
}

// showEntryHandler for the "GET /v1/entries/:id" endpoint
func (app *application) showRandStringHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Display the random string
	num := int(id) //Convert int64 id to int
	fmt.Fprintf(w, "show random string for %d\n", num)
	fmt.Fprintln(w, app.Tools.generateRandomString(num))
}

// showSchoolHandler for the "GET /v1/user/:id" endpoint
func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Create a new instance of the School struct containing the ID we extracted
	// from our URL and some sample data
	school := data.User{
		ID:         id,
		CreatedAt:  time.Now(),
		Name:       "Rene Sanchez",
		Phone:      "623-5411",
		Address:    "20 Caterpillar street",
		Occupation: "Student",
		Version:    1,
	}
	err = app.writeJSON(w, http.StatusOK, school, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
