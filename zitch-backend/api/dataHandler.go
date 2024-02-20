package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Data struct {
	ID        uuid.UUID
	ImgUrl    string    `json:"imgurl"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}

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
		ImgUrl:    input.ImgUrl,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Timestamp: time,
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) readDataHandler(w http.ResponseWriter, r *http.Request) {

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
	
}

func (app *application) deleteDataHandler(w http.ResponseWriter, r *http.Request) {

}
