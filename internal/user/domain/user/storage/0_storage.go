package storage

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
)

type userStorage struct {
}

func NewUserStorage() user.Storage {
	return &userStorage{}
}
