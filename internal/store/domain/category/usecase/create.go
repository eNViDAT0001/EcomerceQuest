package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage/io"
)

func (u *categoryUseCase) CreateCategory(ctx context.Context, input *ioSto.CategoryForm) error {
	return u.categorySto.CreateCategory(ctx, input)
}
