package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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

	docID := createRecord(client, ctx)
	readRecord(client, ctx, docID)
	searchRecords(client, ctx, "Ada", "Lovelace", 1815)
}

func createRecord(client *firestore.Client, ctx context.Context) string {
	docRef, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding a new record: %v", err)
	}
	fmt.Println("Record added successfully")
	return docRef.ID
}

func readRecord(client *firestore.Client, ctx context.Context, docID string) {
	doc, err := client.Collection("users").Doc(docID).Get(ctx)
	if err != nil {
		log.Fatalf("Failed to get document: %v", err)
	}
	fmt.Printf("Retrieved document data: %#v\n", doc.Data())
}

func searchRecords(client *firestore.Client, ctx context.Context, first, last string, born int) {
	iter := client.Collection("users").
		Where("first", "==", first).
		Where("last", "==", last).
		Where("born", "==", born).
		Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Printf("Document found: %v\n", doc.Data())
	}
}
