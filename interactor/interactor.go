package interactor

import (
	"github.com/a2-ito/go-echo-onion-sample/presentation/http/echo/handler"
)

type (
	Interactor interface {
		NewAppHandler() handler.AppHandler
		NewUserHandler() handler.UserHandler
	}

	interactor struct {
	}
)

func NewInteractor() Interactor {
	return nil
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	//appHandler.UserHandler = i.NewUserHandler()
	return appHandler
}
