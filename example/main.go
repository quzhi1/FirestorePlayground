package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8081")

	// Initialize the Firestore client
	client, err := firestore.NewClient(ctx, "my-project", option.WithoutAuthentication())
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	fmt.Println("Firestore client initialized successfully")
}
