package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/ryohei1216/firebase-learn/domain/entity/user"
	"github.com/ryohei1216/firebase-learn/domain/repository"
)

type userRepository struct {
	sc *firestore.Client
}

func NewUserRepository(storeClient *firestore.Client) repository.UserRepository {
	return &userRepository{
		sc: storeClient,
	}
}

func (ur userRepository) Create(ctx context.Context, u user.User) (*user.User, error) {
	ref, result, err := ur.sc.Collection("users").Add(ctx, u)
	if err != nil {
		log.Printf("failed to create user: %v", err)
	}
	log.Println("ref", ref)
	log.Println("result", result)

	return nil, nil
}
