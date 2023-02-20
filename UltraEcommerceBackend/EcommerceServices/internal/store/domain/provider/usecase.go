package provider

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/provider/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

type UseCase interface {
	CreateProvider(ctx context.Context, input io.ProviderForm) (io.ProviderForm, error)
	GetProviderByID(ctx context.Context, userID uint) (entities.Provider, error)
	UpdateProvider(ctx context.Context, providerID uint, input io.ProviderUpdateForm) error
	DeleteProviderByIDs(ctx context.Context, providerID []uint) error
	ListProviderByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) (providers []entities.Provider, total int64, err error)
	ListProvider(ctx context.Context) ([]entities.Provider, error)
}
