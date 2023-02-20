package category

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/usecase/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

type UseCase interface {
	ListCategories(ctx context.Context) ([]entities.Category, error)
	CreateCategory(ctx context.Context, input *ioSto.CategoryForm) error
	UpdateCategory(ctx context.Context, categoryID uint, input *ioSto.CategoryForm) error
	GetCategoryChildrenTreeWithCategoryID(ctx context.Context, categoryID uint) (ioSto.CategoryChildrenTree, error)
	GetCategoryParentsTreeWithCategoryID(ctx context.Context, categoryID uint) (ioSto.CategoryChildrenTree, error)
	GetCategoryChildrenByCategoryID(ctx context.Context, categoryID uint) ([]entities.Category, error)
	GetCategoryRoofList(ctx context.Context) ([]ioUC.CategoryPreview, error)
	GetCategoryTree(ctx context.Context) ([]ioSto.CategoryChildrenTree, error)
}
