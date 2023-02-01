package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

type categoryUseCase struct {
	categorySto category.Storage
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
