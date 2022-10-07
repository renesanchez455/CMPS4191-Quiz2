/*
	CMPS4191 - Quiz #1
	Rene Sanchez - 2018118383
*/
// Filename: cmd/api/routes.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Create a new httprouter instance
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/files", app.uploadFileHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/rand/:id", app.showRandStringHandler)

	return router
}