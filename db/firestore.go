// db/firestore.go
package db

import (
	"context"

	"cloud.google.com/go/firestore"
)

var FirestoreClient *firestore.Client

func InitFirestore() error {
   ctx := context.Background()

   // Replace "your-project-id" with your actual Firebase project ID
   client, err := firestore.NewClient(ctx, "gsc-iiit-kota")
   if err != nil {
      return err
   }

   FirestoreClient = client
   return nil
}
