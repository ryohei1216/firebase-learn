package usecase

import (
	"context"

	"github.com/ryohei1216/firebase-learn/domain/entity/user"
	"github.com/ryohei1216/firebase-learn/domain/repository"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, email string, password string) (*user.User, error)
	GetUser(ctx context.Context, uid string) (*user.User, error)
	UpdateUser(ctx context.Context, uid string, email string, password string) (*user.User, error)
	DeleteUser(ctx context.Context, uid string) error
}

type userUsecase struct {
	urr repository.UserRecordRepository
	ur  repository.UserRepository
}

func NewUserUsecase(urr repository.UserRecordRepository, ur  repository.UserRepository) UserUsecase {
	return &userUsecase{
		urr: urr,
		ur:  ur,
	}
}

func (uu userUsecase) CreateUser(ctx context.Context, email string, password string) (*user.User, error) {
	ur, err := uu.urr.Create(ctx, email, password)
	if err != nil {
		return nil, err
	}

	id, err := user.NewID(ur.UID)
	if err != nil {
		return nil, err
	}

	e, err := user.NewEmail(email)
	if err != nil {
		return nil, err
	}

	p, err := user.NewPassword(password)
	if err != nil {
		return nil, err
	}

	// これをfirestoreにつめる
	u, err := user.New(id, e, p)
	if err != nil {
		return nil, err
	}

	u.SetUserRecord(ur)

	_, err = uu.ur.Create(ctx, *u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu userUsecase) GetUser(ctx context.Context, uid string) (*user.User, error) {
	u, err := uu.ur.Get(ctx, uid)
	if err != nil {
		return nil, err
	}

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

	u, err := user.New("", e, p)
	if err != nil {
		return nil, err
	}
	ur, err := uu.urr.Update(ctx, uid, u)
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
	err := uu.urr.Delete(ctx, uid)
	if err != nil {
		return nil
	}

	return nil
}
