package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/order/domain/order/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/order/entities"
)

func (s orderStorage) ListPreviewByUserID(ctx context.Context, userID uint, input paging.ParamsInput) ([]io.OrderPreview, error) {
	result := make([]io.OrderPreview, 0)
	db := wrap_gorm.GetDB()
	query := db.Table(entities.Order{}.TableName()).
		Select("`Order`.*, JSON_ARRAYAGG(JSON_OBJECT('id', OrderItem.id, 'name', OrderItem.name, 'image', OrderItem.image,'price', OrderItem.price,'discount', OrderItem.discount,'quantity', OrderItem.quantity,'option', OrderItem.option)) as items").
		Joins("JOIN OrderItem on Order.id = OrderItem.order_id").
		Where("Order.user_id = ?", userID).
		Where("Order.deleted_at IS NULL").
		Where("OrderItem.deleted_at IS NULL").
		Group("Order.id")

	paging_query.SetPagingQuery(&input, entities.Order{}.TableName(), query)
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
