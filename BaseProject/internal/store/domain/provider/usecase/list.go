package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (u providerUseCase) ListProvider(ctx context.Context) ([]entities.Provider, error) {
	return u.providerSto.ListProvider(ctx)
}
