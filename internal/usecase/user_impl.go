package usecase

import (
	"context"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/model"
)

type user struct {
}

func NewUser() User {
	return &user{}
}

func (u *user) GetUser(ctx context.Context, id string) (*model.User, error) {
	return &model.User{
		ID:    id,
		Name:  "UserName",
		Email: "example.com",
	}, nil
}
