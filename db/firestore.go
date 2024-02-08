// db/firestore.go
package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var FirestoreClient *firestore.Client
var AuthClient *auth.Client

func InitFirestore() error {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "gsc-iiit-kota"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth client: %v\n", err)
	}

	FirestoreClient = client
	AuthClient = authClient

	return nil
}
