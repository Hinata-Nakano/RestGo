package user

import (
	"example.com/RestCRUD/domain"
	"example.com/RestCRUD/domain/user"
)

type CreateUserUseCase struct {
	repo  user.UserRepository
	clock domain.Clock
}

func NewCreateUserUseCase(repo user.UserRepository, clock domain.Clock) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo, clock: clock}
}

func (u *CreateUserUseCase) Execute(request user.CreateUserRequest) (*user.User, error) {
	userEntity, err := user.NewUser(request, u.clock)
	if err != nil {
		return nil, err
	}

	return u.repo.Create(userEntity)
}
