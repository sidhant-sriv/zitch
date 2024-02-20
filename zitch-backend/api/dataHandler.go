package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (app *application) createDataHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ImgUrl    string  `json:"imgurl"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Timestamp string  `json:"timestamp"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	time, err := time.Parse("2006-01-02 15:04", input.Timestamp)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	data := &Data{
		ID:        uuid.New(),
		ImgUrl:    input.ImgUrl,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Timestamp: time,
	}

	if err := data.Validate(); err != nil {
		app.failedValidationResponse(w, r, err)
	}

	// Function to create data here
	err = createData(data)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) readDataHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readUUIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Function to get Data here
	data, err := getData(id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) updateDataHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ImgUrl    *string  `json:"imgurl"`
		Latitude  *float64 `json:"latitude"`
		Longitude *float64 `json:"longitude"`
		Timestamp *string  `json:"timestamp"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	id, err := app.readUUIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Function to get Data here
	data, err := getData(id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if input.ImgUrl != nil {
		data.ImgUrl = *input.ImgUrl
	}
	if input.Latitude != nil {
		data.Latitude = *input.Latitude
	}
	if input.Longitude != nil {
		data.Longitude = *input.Longitude
	}
	if input.Timestamp != nil {
		data.Timestamp, err = time.Parse("2006-01-02 15:04", *input.Timestamp)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	if err := data.Validate(); err != nil {
		app.failedValidationResponse(w, r, err)
	}

	// Function to update Data here
	err = updateData(data)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) deleteDataHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readUUIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Function to delete Data here
	err = deleteData(id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "Data successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
