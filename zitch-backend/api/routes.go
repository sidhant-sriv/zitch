package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/data", app.createDataHandler)
	router.HandlerFunc(http.MethodGet, "/v1/data/:id", app.readDataHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/data/:id", app.updateDataHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/data/:id", app.deleteDataHandler)

	return router
}
