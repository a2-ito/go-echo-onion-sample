package interactor

import (
  "github.com/a2-ito/go-echo-onion-sample/presentation/http/echo/handler"
)

type {
  Interactor interface {
    NewUserHnadler() handler.UserHnadler
  }

  interactor struct {
  }
}

func NewInteractor() Interactor {
  return &interactor{}
}

func (i *interactor) NewUserHandler() handler.UserHnadler {
  return handler.NewUserHnadler(i.NewUserUseCAse())
}
