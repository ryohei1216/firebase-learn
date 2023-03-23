package usecase

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/ryohei1216/firebase-learn/repository"
)

type UserUsecase interface {
	GetUser(ctx context.Context, uid string) (*auth.UserRecord, error)
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		ur: ur,
	}
}

func (uu userUsecase) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	u, err := uu.ur.Get(ctx, uid)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, err
	}

	return u, nil
}