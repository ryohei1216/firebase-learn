package repository

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

type UserRepository interface {
	Get(ctx context.Context, uid string) (*auth.UserRecord, error)
}