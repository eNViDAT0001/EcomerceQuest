package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/storage/io"
)

func (u *categoryUseCase) GetCategoryTree(ctx context.Context) ([]ioSto.CategoryChildrenTree, error) {
	categoryRoofTree := make([]ioSto.CategoryChildrenTree, 0)
	//Get Roof Category
	roofCategories, err := u.categorySto.GetCategoryRoofTree(ctx)
	if err != nil {
		return nil, err
	}
	//Get tree from roof
	for _, category := range roofCategories {
		tree, err := u.GetCategoryChildrenTreeWithCategoryID(ctx, category.ID)
		if err != nil {
			return nil, err
		}
		tree.ID = category.ID
		tree.Name = category.Name
		tree.ImagePath = category.ImagePath
		categoryRoofTree = append(categoryRoofTree, tree)
	}

	return categoryRoofTree, nil
}
