package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
	"gorm.io/gorm"
)

func (s productStorage) GetSpecificationTreeByProductID(ctx context.Context, productID uint) ([]io.ProductSpecificationFullDetail, error) {

	result := make([]io.ProductSpecificationFullDetail, 0)
	db := wrap_gorm.GetDB()
	query := db.Raw("WITH recursive cte (id, product_id, properties) AS (SELECT id, product_id, properties FROM ProductSpecification WHERE ProductSpecification.specification_id IS NULL AND ProductSpecification.product_id = ? UNION ALL SELECT p.id, p.product_id, p.properties FROM ProductSpecification p INNER JOIN cte ON p.specification_id = cte.id) SELECT cte.id, cte.properties, IF(COUNT(ProductOption.id) = 0, NULL, JSON_ARRAYAGG(JSON_OBJECT( 'id', ProductOption.id, 'name', ProductOption.name, 'price', ProductOption.price, 'quantity', ProductOption.quantity))) AS options FROM cte INNER JOIN ProductOption ON ProductOption.specification_id = cte.id GROUP BY cte.id, cte.properties;", productID).
		Scan(&result)

	err := query.Error
	if err != nil {
		return result, err
	}
	if query.RowsAffected < 1 {
		return result, gorm.ErrRecordNotFound
	}
	return result, err
}
