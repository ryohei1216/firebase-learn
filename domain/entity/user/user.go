package user

import "firebase.google.com/go/v4/auth"

type User struct {
	Email      Email
	Password   Password
	UserRecord *auth.UserRecord
}

func New(email string, password string) (*User, error) {
	e, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	p, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:    e,
		Password: p,
	}, nil
}

func (u *User) SetUserRecord(ur *auth.UserRecord) error {
	u.UserRecord = ur
	return nil
}