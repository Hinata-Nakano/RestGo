package user

import (
	"time"

	"example.com/RestCRUD/domain"
	"github.com/google/uuid"
)

type UserID string
type UserName string
type UserEmail string
type UserPassword string

type CreateUserRequest struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func newUserName(name string) (UserName, error) {
	if name == "" {
		return "", domain.ErrNameRequired
	}
	return UserName(name), nil
}

func newUserEmail(email string) (UserEmail, error) {
	if email == "" {
		return "", domain.ErrEmailRequired
	}
	return UserEmail(email), nil
}

func newUserPassword(password string) UserPassword {
	return UserPassword(password)
}

func newUserID() UserID {
	return UserID(uuid.New().String())
}

type User struct {
	ID        UserID
	Name      UserName
	Email     UserEmail
	Password  UserPassword
	CreatedAt string
	UpdatedAt string
}

func NewUser(createUserRequest CreateUserRequest, clock domain.Clock) (*User, error) {
	name, err := newUserName(createUserRequest.Name)
	if err != nil {
		return nil, err
	}

	email, err := newUserEmail(createUserRequest.Email)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        newUserID(),
		Name:      name,
		Email:     email,
		Password:  newUserPassword(createUserRequest.Password),
		CreatedAt: clock.Now().Format(time.RFC3339),
		UpdatedAt: clock.Now().Format(time.RFC3339),
	}, nil
}
