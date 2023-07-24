package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

type categoryUseCase struct {
	categorySto category.Storage
}

func (u *categoryUseCase) DeleteCategoryByID(ctx context.Context, categoryID uint, parentID *uint) error {
	return u.categorySto.DeleteCategoryByID(ctx, categoryID, parentID)
}

func (u *categoryUseCase) DeleteCategoryNodeByID(ctx context.Context, categoryID uint) error {
	return u.categorySto.DeleteCategoryNodeByID(ctx, categoryID)
}

func (u *categoryUseCase) RecoverCategoryByID(ctx context.Context, categoryID uint) error {
	return u.categorySto.RecoverCategoryByID(ctx, categoryID)
}

func (u *categoryUseCase) RecoverCategoryNodeByID(ctx context.Context, categoryID uint) error {
	return u.RecoverCategoryNodeByID(ctx, categoryID)
}

func (u *categoryUseCase) ListCategories(ctx context.Context) ([]entities.Category, error) {
	return u.categorySto.ListCategories(ctx)
}

func NewCategoryUseCase(
	categorySto category.Storage,

) category.UseCase {
	return &categoryUseCase{
		categorySto: categorySto,
	}
}
