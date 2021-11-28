package datastore

import (
  "context"

  "github.com/a2-ito/go-echo-onion-sample/domain/repository"
  "github.com/jinzhu/gorm"
)

type userRepository struct {
  Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) repository.UserRepository {
  return &userRepository{Conn}
}

func (r *userRepository) Fetch(ctx context.Context) error {
  return nil
}
