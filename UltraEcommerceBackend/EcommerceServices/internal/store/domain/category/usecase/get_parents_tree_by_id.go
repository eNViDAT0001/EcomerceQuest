package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/storage"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/usecase/convert"
)

func (u *categoryUseCase) GetCategoryParentsTreeWithCategoryID(ctx context.Context, categoryID uint) (ioSto.CategoryChildrenTree, error) {
	var result ioSto.CategoryChildrenTree
	categories, err := u.categorySto.GetCategoryParentsTreeWithCategoryID(ctx, categoryID)
	if err != nil {

		return result, err
	}

	if len(categories) == 0 {
		return result, nil
	}

	categoriesTree, err := convert.ArrayCategoryEntityToCategoryChildrenTree(categories)
	if err != nil {
		return result, err
	}

	result = categoriesTree[len(categoriesTree)-1]
	for i := len(categoriesTree) - 2; i >= 0; i-- {
		storage.GetCategoryParentTree(&result, categoriesTree[i])
	}

	return result, nil
}
