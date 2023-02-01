package category

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

type Storage interface {
	ListCategories(ctx context.Context) ([]entities.Category, error)
	CreateCategory(ctx context.Context, input *ioSto.CategoryForm) error
	UpdateCategory(ctx context.Context, categoryID uint, input *ioSto.CategoryForm) error
	GetCategoryChildrenTreeWithCategoryID(ctx context.Context, categoryID uint) ([]entities.Category, error)
	GetCategoryParentsTreeWithCategoryID(ctx context.Context, categoryID uint) ([]entities.Category, error)
	GetCategoryRoofTree(ctx context.Context) ([]entities.Category, error)
}
