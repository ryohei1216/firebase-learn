package repository

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
)

type UserRepository interface {
	Get(ctx context.Context, uid string) (*auth.UserRecord, error)
}

type userRepository struct {
	firebaseClient *auth.Client
}

func NewUserRepository(firebaseClient *auth.Client) UserRepository {
	return &userRepository{
		firebaseClient: firebaseClient,
	}
}

func (ur userRepository) Get(ctx context.Context, uid string) (*auth.UserRecord, error) {
	u, err := ur.firebaseClient.GetUser(ctx, uid)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, err
	}

	return u, nil
}