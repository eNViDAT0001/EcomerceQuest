package convert

import (
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"github.com/jinzhu/copier"
)

func ArrayCategoryEntityToCategoryChildrenTree(inputEntity []entities.Category) ([]ioSto.CategoryChildrenTree, error) {
	result := make([]ioSto.CategoryChildrenTree, 0)

	err := copier.Copy(&result, &inputEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func ArrayCategoryEntityToCategoryParentTree(inputEntity []entities.Category) ([]ioSto.CategoryParentTree, error) {
	result := make([]ioSto.CategoryParentTree, 0)

	err := copier.Copy(&result, &inputEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}
