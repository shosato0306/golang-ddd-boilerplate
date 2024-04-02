package repository

import (
	"context"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/model"
)

type User interface {
	GetUserByID(
		ctx context.Context,
		id string,
	) (*model.User, error)
	Create(
		ctx context.Context,
		user *model.User,
	) (*model.User, error)
}
