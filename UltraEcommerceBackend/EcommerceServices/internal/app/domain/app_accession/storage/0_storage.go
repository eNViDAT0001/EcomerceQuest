package storage

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_viper"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/app/domain/app_accession"
)

var wViper = wrap_viper.GetViper()

type appAccessionStorage struct {
}

func NewAppAccessionStorage() app_accession.Storage {
	return &appAccessionStorage{}
}
