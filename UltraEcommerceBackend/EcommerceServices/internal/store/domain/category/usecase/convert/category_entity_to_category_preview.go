package convert

import (
	ioUC "github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/usecase/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
	"github.com/jinzhu/copier"
)

func CategoryEntityToCategoryPreview(inputEntity *entities.Category) (*ioUC.CategoryPreview, error) {
	var result ioUC.CategoryPreview
	err := copier.Copy(&result, inputEntity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func ArrayCategoryEntityToCategoryPreview(inputEntity []entities.Category) ([]ioUC.CategoryPreview, error) {
	result := make([]ioUC.CategoryPreview, 0)

	err := copier.Copy(&result, &inputEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}
