package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"gorm.io/gorm"
)

func (u *categoryUseCase) UpdateCategory(ctx context.Context, categoryID uint, input *ioSto.CategoryForm) error {
	if input.CategoryParentID == nil {
		return u.categorySto.UpdateCategory(ctx, categoryID, input)
	}

	children, err := u.categorySto.GetCategoryChildrenTreeWithCategoryID(ctx, categoryID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.categorySto.UpdateCategory(ctx, categoryID, input)
		}
		return err
	}

	var parentDetail entities.Category
	var childDetail *entities.Category
	for _, child := range children {
		if *child.CategoryParentID == categoryID {
			childDetail = &child
			parentDetail, err = u.categorySto.GetCategoryDetailByID(ctx, categoryID)
			if err != nil {
				return err
			}
			break
		}
	}

	if childDetail == nil {
		newChild := ioSto.CategoryForm{
			CategoryParentID: parentDetail.CategoryParentID,
		}
		err = u.categorySto.UpdateCategory(ctx, childDetail.ID, &newChild)
		if err != nil {
			return err
		}
	}

	return u.categorySto.UpdateCategory(ctx, categoryID, input)
}
