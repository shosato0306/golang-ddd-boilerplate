package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/model"
	"github.com/shosato0306/golang-ddd-boilerplate/internal/domain/repository"
	"github.com/shosato0306/golang-ddd-boilerplate/internal/infrastructure/repository/dbconn"
	"github.com/shosato0306/golang-ddd-boilerplate/internal/infrastructure/repository/dbmodels"
)

type user struct {
}

func NewUser() repository.User {
	return &user{}
}

func (r *user) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	mods := []qm.QueryMod{}
	mods = append(mods, dbmodels.UserWhere.ID.EQ(id))

	dbUser, err := dbmodels.Users(mods...).
		One(ctx, dbconn.GetContextExecutor(ctx))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
	}

	return &model.User{
		ID:    dbUser.ID,
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}, nil
}

func (r *user) Create(
	ctx context.Context,
	user *model.User,
) (*model.User, error) {
	dbUser := dbmodels.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	err := dbUser.Insert(
		ctx, dbconn.GetContextExecutor(ctx), boil.Infer())
	if err != nil {
		return nil, err
	}

	return user, nil
}
