package usecase

import (
	"context"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/model"
)

type User interface {
	GetUser(
		ctx context.Context,
		id string,
	) (*model.User, error)
	CreateUser(
		ctx context.Context,
		name string,
		email string,
	) (*model.User, error)
}
