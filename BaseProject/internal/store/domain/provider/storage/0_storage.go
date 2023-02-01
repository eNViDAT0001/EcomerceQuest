package storage

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider"
)

type providerStorage struct {
}

func NewProviderStorage() provider.Storage {
	return &providerStorage{}
}
