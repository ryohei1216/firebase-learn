package user

type Password string

func NewPassword(password string) (Password, error) {
	return Password(password), nil
}