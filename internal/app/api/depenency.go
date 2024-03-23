package api

import (
	"github.com/shosato0306/golang-ddd-boilerplate/internal/app/api/handler"
	"github.com/shosato0306/golang-ddd-boilerplate/internal/usecase"
)

type Dependency struct {
	exampleHandler *handler.ExampleHandler
}

func (d *Dependency) Inject() {
	userUsecase := usecase.NewUser()

	d.exampleHandler = handler.NewExampleHandler(userUsecase)
}
