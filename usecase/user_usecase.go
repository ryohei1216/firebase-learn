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
	DeleteUser(ctx context.Context, uid string) error
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
	e, err := user.NewEmail(email)
	if err != nil {
		return nil, err
	}

	p, err := user.NewPassword(password)
	if err != nil {
		return nil, err
	}
	
	u, err := user.New(e, p)
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

func (uu userUsecase) DeleteUser(ctx context.Context, uid string) error {
	err := uu.ur.Delete(ctx, uid)
	if err != nil {
		return nil
	}

	return nil
}
