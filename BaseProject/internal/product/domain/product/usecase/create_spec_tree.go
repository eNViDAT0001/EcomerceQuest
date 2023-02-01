package usecase

import (
	"context"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/io"
)

func (u *productUseCase) CreateSpecificationTree(ctx context.Context, specs []ioUC.SpecificationCreateForm) error {
	specID, err := u.CreateSpecification(ctx, specs[0])
	if err != nil {
		return err
	}

	for i := 1; i < len(specs); i++ {
		specs[i].Specification.SpecificationID = &specID
		id, err := u.CreateSpecification(ctx, specs[i])
		if err != nil {
			return err
		}
		specID = id
	}
	return nil
}
