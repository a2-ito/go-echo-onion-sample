package usecase

import (
	"context"
	"github.com/a2-ito/go-echo-onion-sample/domain/repository"
)

type (
	UserUseCase interface {
		Get(ctx context.Context) error
		GetUsers(ctx context.Context) error
	}
)

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) Get(ctx context.Context) error {
	return u.UserRepository.Fetch(ctx)
}

func (u *userUseCase) GetUsers(ctx context.Context) error {
	return u.UserRepository.Fetch(ctx)
}
