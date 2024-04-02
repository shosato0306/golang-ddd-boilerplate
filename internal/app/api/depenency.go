package api

import (
	"github.com/shosato0306/golang-ddd-boilerplate/internal/app/api/handler"
	"github.com/shosato0306/golang-ddd-boilerplate/internal/infrastructure/repository"
	"github.com/shosato0306/golang-ddd-boilerplate/internal/usecase"
)

type Dependency struct {
	apiHandler *handler.APIHandler
}

func (d *Dependency) Inject() {
	userRepository := repository.NewUser()

	userUsecase := usecase.NewUser(userRepository)

	d.apiHandler = handler.NewAPIHandler(userUsecase)
}
