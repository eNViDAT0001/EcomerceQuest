package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/entities"
)

func (a addressUseCase) GetProvinces(ctx context.Context) ([]entities.Province, error) {
	return a.addressSto.GetProvinces(ctx)
}
