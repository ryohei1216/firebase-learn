package user

import "firebase.google.com/go/v4/auth"

type User struct {
	Email      Email
	Password   Password
	UserRecord *auth.UserRecord
}

func New(email Email, password Password) (*User, error) {
	return &User{
		Email:    email,
		Password: password,
	}, nil
}

func (u *User) SetUserRecord(ur *auth.UserRecord) error {
	u.UserRecord = ur
	return nil
}
