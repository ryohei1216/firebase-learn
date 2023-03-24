package repository

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/ryohei1216/firebase-learn/domain/entity/user"
)

type UserRepository interface {
	Get(ctx context.Context, uid string) (*auth.UserRecord, error)
	Update(ctx context.Context, uid string, user *user.User) (*auth.UserRecord, error)
}