package services

import "github.com/Resoluter/bookstore-users.git/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
