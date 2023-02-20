package entities

type ProductSpecification struct {
	ID              uint   `gorm:"primarykey" json:"id"`
	SpecificationID *uint  `gorm:"column:specification_id" json:"specification_id"`
	ProductID       uint   `gorm:"column:product_id" json:"product_id"`
	Properties      string `gorm:"column:properties" json:"properties"`
}

func (ProductSpecification) TableName() string {
	return "ProductSpecification"
}
