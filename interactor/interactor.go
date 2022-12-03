package interactor

import (
	"github.com/a2-ito/go-echo-onion-sample/application"
	//"github.com/a2-ito/go-echo-onion-sample/domain/repository"
	repository "repository"
	"github.com/a2-ito/go-echo-onion-sample/infrastructure/persistence/datastore"
	//"github.com/a2-ito/go-echo-onion-sample/presentation/http/echo/handler"
	handler "handler"
	//"github.com/jinzhu/gorm"
)

type (
	Interactor interface {
		NewAppHandler() handler.AppHandler
		NewUserHandler() handler.UserHandler
		NewUserUseCase() usecase.UserUseCase
		NewUserRepository() repository.UserRepository
	}

	interactor struct {
		//Conn *gorm.DB
	}
)

/*
	func NewInteractor(Conn *gorm.DB) Interactor {
		return &interactor{Conn}
	}
*/
func NewInteractor() Interactor {
	return &interactor{}
}

type appHandler struct {
	handler.UserHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	return appHandler
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	//return datastore.NewUserRepository(i.Conn)
	return datastore.NewUserRepository()
}
