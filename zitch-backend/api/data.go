package main

import (
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

func (data *Data) Validate() map[string]string {
	Errors := make(map[string]string)

	if !(-90 <= data.Latitude || data.Latitude <= 90) {
		if _, exists := Errors["latitude"]; !exists {
			Errors["latitude"] = "must be a valid latitude"
		}
	}

	if !(-180 <= data.Longitude || data.Longitude <= 180) {
		if _, exists := Errors["longitude"]; !exists {
			Errors["longitude"] = "must be a valid longitude"
		}
	}

	return Errors
}

func createData(data *Data) error {
	return nil
}

func getData(id uuid.UUID) (*Data, error) {
	return nil, nil
}

func updateData(data *Data) error {
	return nil
}

func deleteData(id uuid.UUID) error {
	return nil
}
