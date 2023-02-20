package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/provider/storage/io"
)

func (u providerUseCase) CreateProvider(ctx context.Context, input io.ProviderForm) (io.ProviderForm, error) {
	return u.providerSto.CreateProvider(ctx, input)
}
