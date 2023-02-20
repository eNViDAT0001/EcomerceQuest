package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/usecase/convert"
	ioUC "github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/usecase/io"
)

func (u *categoryUseCase) GetCategoryRoofList(ctx context.Context) ([]ioUC.CategoryPreview, error) {

	categories, err := u.categorySto.GetCategoryRoofTree(ctx)
	if err != nil {
		return nil, err
	}
	result, err := convert.ArrayCategoryEntityToCategoryPreview(categories)
	if err != nil {
		return nil, err
	}

	return result, nil

}
