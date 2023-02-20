package storage

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/domain/user"
)

type userStorage struct {
}

func NewUserStorage() user.Storage {
	return &userStorage{}
}
