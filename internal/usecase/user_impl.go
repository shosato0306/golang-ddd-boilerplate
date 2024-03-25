package usecase

import (
	"context"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/model"
	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/repository"
)

type user struct {
	rUser repository.User
}

func NewUser(
	rUser repository.User,
) User {
	return &user{
		rUser: rUser,
	}
}

func (u *user) GetUser(ctx context.Context, id string) (*model.User, error) {
	user, err := u.rUser.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
