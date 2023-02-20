package storage

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_viper"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/verification/domain/jwt"
)

var wViper = wrap_viper.GetViper()

type jwtStorage struct {
}

func NewJwtStorage() jwt.Storage {
	return &jwtStorage{}
}
