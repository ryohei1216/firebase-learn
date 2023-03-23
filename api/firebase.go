package api

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func NewFirebaseClient() *auth.Client {
	opt := option.WithCredentialsFile("firebase_secret.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println(err)
	}
	
	ctx := context.Background()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Printf("error getting Auth client: %v\n", err)
	}

	return client
}