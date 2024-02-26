package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"zitch/backend/api/inference"
)

func (app *application) callClassifierHandler(w http.ResponseWriter, r *http.Request) {
	// Read the file from the request body
	file, _, err := r.FormFile("image")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	defer file.Close()

	// Read the file data into a byte slice
	fileData, err := io.ReadAll(file)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Convert the byte slice into an image

	tempDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	defer os.RemoveAll(tempDir)

	tempFile, err := os.CreateTemp("/temp", "image_*.jpg")
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	defer tempFile.Close()

	_, err = tempFile.Write(fileData)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	imagePath := tempFile.Name()
	fmt.Println(imagePath)
	result, conf, err := inference.Inference(imagePath)
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
