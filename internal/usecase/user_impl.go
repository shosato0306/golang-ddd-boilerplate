package usecase

import (
	"context"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/model"
	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/repository"
	"github.com/shosato0306/golang-ddd-boilerplate/internal/infrastructure/repository/dbconn"
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

func (u *user) GetUser(
	ctx context.Context,
	id string,
) (*model.User, error) {
	user, err := u.rUser.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) CreateUser(
	ctx context.Context,
	name string,
	email string,
) (*model.User, error) {
	user := model.NewUser(name, email)

	if err := dbconn.RunTx(ctx, func(ctx context.Context) error {
		_, err := u.rUser.Create(ctx, user)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return user, nil
}
