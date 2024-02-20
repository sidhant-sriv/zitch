package main

import (
	"net/http"
)

// healthcheckHandler is a basic request used to check the status of the server
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"version":     version,
			"environment": app.config.env,
		},
	}

	app.writeJSON(w, http.StatusOK, data, nil)
}
