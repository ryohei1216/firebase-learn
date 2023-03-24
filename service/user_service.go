package service

import (
	"context"

	"github.com/ryohei1216/firebase-learn/domain/entity/user"
	"github.com/ryohei1216/firebase-learn/usecase"
)

type UserService interface {
	GetUser(ctx context.Context, uid string) (*user.User, error)
	UpdateUser(ctx context.Context, uid string , email string, password string) (*user.User, error)
}

type userService struct {
	uu usecase.UserUsecase
}

func NewUserService(uu usecase.UserUsecase) UserService {
	return &userService{
		uu: uu,
	}
}

func (us *userService) GetUser(ctx context.Context, uid string) (*user.User, error) {
	u, err := us.uu.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *userService) UpdateUser(ctx context.Context, uid string, email string, password string) (*user.User, error) {
	u, err := us.uu.UpdateUser(ctx, uid, email, password)
	if err != nil {
		return nil, err
	}

	return u, nil
}