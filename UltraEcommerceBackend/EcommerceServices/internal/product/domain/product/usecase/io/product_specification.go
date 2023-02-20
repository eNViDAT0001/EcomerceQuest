package io

import "github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"

type SpecificationCreateForm struct {
	Specification io.ProductSpecificationCreateForm
	Options       []io.ProductOptionCreateForm
}
