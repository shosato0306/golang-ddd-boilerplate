package model

import "github.com/shosato0306/golang-ddd-boilerplate/internal/pkg"

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUser(
	name string,
	email string,
) *User {
	return &User{
		ID:    pkg.NewULID(),
		Name:  name,
		Email: email,
	}
}
