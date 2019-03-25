package repository

import (
	models "go-postgres/model"
)

type UserRepo interface {
	Select() ([]models.User, error)
	Insert(u models.User) (error)
}
