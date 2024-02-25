// package main

// import (
//     "context"
//     "log"
//     "time"

//     // pb "api/inference"

//     "google.golang.org/grpc"
// )

// const (
//     address     = "localhost:50051"
// )

// func makeInference(imagePath string) {
//     // Set up a connection to the server at port 50051.
//     conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
//     if err != nil {
//         log.Fatalf("did not connect: %v", err)
//     }
//     defer conn.Close()
//     c := NewInferenceClient(conn)

//     // Prepare the image data.
//     // Here we're using a placeholder empty byte slice.
//     // You should replace this with your actual image data.
//     imageData := []byte{}

//     // Contact the server and print out its response.
//     ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//     defer cancel()
//     r, err := c.Predict(ctx, &Image{Content: imageData})
//     if err != nil {
//         log.Fatalf("could not predict: %v", err)
//     }
//     log.Printf("Prediction: value = %d, confidence = %d", r.GetValue(), r.GetConfidence())
// }

package inference

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func Inference(imageData []byte) (int32, int32, error) {
	// Read image data
	// imageData, err := os.ReadFile(imagePath)

	// if err != nil {
	// 	return -1, -1, fmt.Errorf("error reading image: %v", err)
	// }
	// Connect to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return -1, -1, fmt.Errorf("error connecting to server: %v", err)
	}
	defer conn.Close()

	// Create the client stub
	client := NewInferenceClient(conn)

	// Start time measurement

	// Send prediction request
	ctx := context.Background()
	response, err := client.Predict(ctx, &Image{Content: imageData})
	if err != nil {
		return -1, -1, fmt.Errorf("error getting prediction: %v", err)
	}

	// Calculate and print execution time
	return response.GetValue(), response.GetConfidence(), nil
}
