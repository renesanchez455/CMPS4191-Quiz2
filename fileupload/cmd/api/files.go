/*
	CMPS4191 - Quiz #2
	Rene Sanchez - 2018118383
*/
// Filename: cmd/api/files.go

package main

import (
	"fmt"
	"net/http"
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
