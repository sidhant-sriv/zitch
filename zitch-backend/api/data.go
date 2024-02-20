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
