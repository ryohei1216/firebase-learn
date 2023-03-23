package service

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/ryohei1216/firebase-learn/usecase"
)

type UserService interface {
	GetUser(ctx context.Context,uid string) (*auth.UserRecord, error)
}

type userService struct {
	uu usecase.UserUsecase
}

func NewUserService(uu usecase.UserUsecase) UserService {
	return &userService{
		uu: uu,
	}
}

func (us *userService) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	u, err := us.uu.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return u, nil
}