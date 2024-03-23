package usecase

import (
	"context"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/model"
)

type User interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
}
