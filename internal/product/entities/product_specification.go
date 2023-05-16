package entities

type ProductSpecification struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	ProductID  uint   `gorm:"column:product_id" json:"product_id"`
	Properties string `gorm:"column:properties" json:"properties"`
}

func (ProductSpecification) TableName() string {
	return "ProductSpecification"
}
