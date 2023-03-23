package infrastructure

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/ryohei1216/firebase-learn/domain/repository"
)

type userRepository struct {
	fc *auth.Client
}

func NewUserRepository(firebaseClient *auth.Client) repository.UserRepository {
	return &userRepository{
		fc: firebaseClient,
	}
}

func (ur userRepository) Get(ctx context.Context, uid string) (*auth.UserRecord, error) {
	u, err := ur.Get(ctx, uid)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, err
	}

	return u, nil
}