package repository

import (
	"context"

	"github.com/ryohei1216/firebase-learn/domain/entity/user"
)

type UserRepository interface {
	Create(ctx context.Context, u user.User) (*user.User, error)
}
