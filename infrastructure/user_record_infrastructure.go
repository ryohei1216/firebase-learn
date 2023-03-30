package infrastructure

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/ryohei1216/firebase-learn/domain/entity/user"
	"github.com/ryohei1216/firebase-learn/domain/repository"
)

type userRecordRepository struct {
	fc *auth.Client
}

func NewUserRepository(firebaseClient *auth.Client) repository.UserRecordRepository {
	return &userRecordRepository{
		fc: firebaseClient,
	}
}

func (ur userRecordRepository) Create(ctx context.Context, email string, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).Email(email).Password(password)
	u, err := ur.fc.CreateUser(ctx, params)
	if err != nil {
		log.Printf("failed to create user: %v", err)
	}

	return u, nil
}

func (ur userRecordRepository) Get(ctx context.Context, uid string) (*auth.UserRecord, error) {
	u, err := ur.fc.GetUser(ctx, uid)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, err
	}

	return u, nil
}

func (ur userRecordRepository) Update(ctx context.Context, uid string, user *user.User) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).Email(string(user.Email)).Password(string(user.Password))
	u, err := ur.fc.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Printf("failed to update user: %v", err)
		return nil, err
	}

	return u, nil
}

func (ur userRecordRepository) Delete(ctx context.Context, uid string) error {
	err := ur.fc.DeleteUser(ctx, uid)
	if err != nil {
		log.Printf("failed to delete user: %v", err)
		return err
	}

	return nil
}
