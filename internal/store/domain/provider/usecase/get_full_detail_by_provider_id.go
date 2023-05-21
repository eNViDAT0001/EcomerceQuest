package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage/io"
)

func (u *providerUseCase) GetProviderFullDetailByID(ctx context.Context, providerID uint) (io.ProviderFullDetail, error) {
	return u.providerSto.GetProviderFullDetailByID(ctx, providerID)
}
