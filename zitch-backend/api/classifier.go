package main

import (
	"net/http"
	"zitch/backend/api/inference"
)

func (app *application) callClassifierHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Imgurl []byte `json:"img"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	result, conf, err := inference.Inference(input.Imgurl)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"result": result, "confidence": conf}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
