package api

import (
	"context"
	"log"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func NewAuthClient() *auth.Client {
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

func AuthMiddleware(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(("Authorization"))
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		_, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(401, gin.H{"message": "invalid id"})
			c.Abort()
		}
		log.Println("Verified ID token:")
	}
}
