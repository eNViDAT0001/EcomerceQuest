package usecase

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage/io"
	"gorm.io/gorm"
)

func (u providerUseCase) ListProviderQuantity(ctx context.Context, filter paging.ParamsInput) (providers []io.ProviderQuantity, total int64, err error) {
	total, err = u.providerSto.CountListProviderQuantity(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	providers, err = u.providerSto.ListProviderQuantity(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return providers, total, err
}
