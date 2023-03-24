package user

type Email string

func NewEmail(email string) (Email, error) {
	return Email(email), nil
}