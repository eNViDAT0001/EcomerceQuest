package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (u providerUseCase) GetProviderByID(ctx context.Context, providerID uint) (entities.Provider, error) {
	return u.providerSto.GetProviderByID(ctx, providerID)
}
func (u providerUseCase) GetProviderFullDetail(ctx context.Context, providerID uint) (entities.Provider, error) {
	return u.providerSto.GetProviderByID(ctx, providerID)
}
