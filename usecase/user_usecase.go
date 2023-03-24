package usecase

import (
	"context"
	"log"

	"github.com/ryohei1216/firebase-learn/domain/entity/user"
	"github.com/ryohei1216/firebase-learn/domain/repository"
)

type UserUsecase interface {
	GetUser(ctx context.Context, uid string) (*user.User, error)
	UpdateUser(ctx context.Context, uid string, email string, password string) (*user.User, error)
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		ur: ur,
	}
}

func (uu userUsecase) GetUser(ctx context.Context, uid string) (*user.User, error) {
	ur, err := uu.ur.Get(ctx, uid)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, err
	}

	u, err := user.New("", "")
	if err != nil {
		return nil, err
	}
	u.SetUserRecord(ur)

	return u, nil
}

func (uu userUsecase) UpdateUser(ctx context.Context, uid string, email string, password string) (*user.User, error) {
	u, err := user.New(email, password)
	if err != nil {
		return nil, err
	}
	ur, err := uu.ur.Update(ctx, uid, u)
	if err != nil {
		return nil, err
	}

	err = u.SetUserRecord(ur)
	if err != nil {
		return nil, err
	}

	return u, nil
}