package storage

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category"
)

type categoryStorage struct {
}

func NewCategoryStorage() category.Storage {
	return &categoryStorage{}
}
