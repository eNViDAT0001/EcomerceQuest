package usecase

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/usecase/convert"
	"gorm.io/gorm"
)

func (u *categoryUseCase) GetCategoryChildrenTreeWithCategoryID(ctx context.Context, categoryID uint) (ioSto.CategoryChildrenTree, error) {
	var result ioSto.CategoryChildrenTree

	baseCategory, err := u.categorySto.GetCategoryDetailByID(ctx, categoryID)
	if err != nil {
		return result, err
	}
	result.ID = categoryID
	result.Name = baseCategory.Name
	result.ImagePath = baseCategory.ImagePath
	if baseCategory.CategoryParentID != nil {
		result.CategoryParentID = *baseCategory.CategoryParentID
	}

	categories, err := u.categorySto.GetCategoryChildrenTreeWithCategoryID(ctx, categoryID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return result, err
	}
	if len(categories) == 0 {
		return result, nil
	}
	categoriesTree, err := convert.ArrayCategoryEntityToCategoryChildrenTree(categories)
	if err != nil {
		return result, err
	}

	for _, v := range categoriesTree {
		storage.GetCategoryChildrenTree(&result, v)
	}

	return result, nil
}
