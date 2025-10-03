package database

import "github.com/PauloPimentel-github/go-expert/APIs/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
