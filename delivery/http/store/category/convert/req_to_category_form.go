package convert

import (
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/category/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage/io"
	"github.com/jinzhu/copier"
)

func UpdateCategoryReqToCategoryForm(inputEntity *io.UpdateCategoryReq) (*ioSto.CategoryForm, error) {
	var result ioSto.CategoryForm
	err := copier.Copy(&result, inputEntity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func CreateCategoryReqToCategoryForm(inputEntity *io.CreateCategoryReq) (*ioSto.CategoryForm, error) {
	var result ioSto.CategoryForm
	err := copier.Copy(&result, inputEntity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
