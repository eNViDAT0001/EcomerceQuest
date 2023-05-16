package storage

import (
	"context"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/fake_redis"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	entities3 "github.com/eNViDAT0001/Thesis/Backend/internal/cart/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	entities2 "github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

func (s orderStorage) CreateOrder(ctx context.Context, order io.CreateOrderForm, items []io2.CreateOrderItemForm, cartItemsIDs []uint) (createdOrders []io.CreateOrderForm, err error) {
	db := wrap_gorm.GetDB()

	providersItems := map[uint][]io2.CreateOrderItemForm{}
	for _, item := range items {
		_, ok := providersItems[item.ProviderID]
		if !ok {
			storage := []io2.CreateOrderItemForm{item}
			providersItems[item.ProviderID] = storage
			continue
		}
		providersItems[item.ProviderID] = append(providersItems[item.ProviderID], item)
	}

	orders := make([]io.CreateOrderForm, 0)
	for k, v := range providersItems {
		quantity := 0
		total := 0.0
		for _, item := range v {
			quantity += item.Quantity
			total += float64(item.Price - item.Discount*(item.Price/100))
		}
		newOrder := io.CreateOrderForm{
			ID:                0,
			UserID:            order.UserID,
			ProviderID:        k,
			Name:              order.Name,
			Gender:            order.Gender,
			Phone:             order.Phone,
			Province:          order.Province,
			District:          order.District,
			Ward:              order.Ward,
			Street:            order.Street,
			Quantity:          quantity,
			Total:             int(total),
			Discount:          order.Discount,
			StatusDescription: order.StatusDescription,
		}
		orders = append(orders, newOrder)
	}

	query := db.Begin()
	// Delete all cart items
	err = query.Table(entities3.CartItem{}.TableName()).
		Where("id IN ?", cartItemsIDs).
		Where("user_id = ?", order.UserID).
		Delete(&entities3.CartItem{}).Error
	if err != nil {
		query.Rollback()
		return nil, err
	}

	err = query.Table(entities.Order{}.TableName()).Create(&orders).Error
	if err != nil {
		query.Rollback()
		return nil, err
	}

	quantityStore := fake_redis.GetQuantityStore()
	store := map[uint]int{}

	//TODO: Fix right now
	for i, v := range items {
		for _, createdOrder := range orders {
			if v.ProviderID == items[i].ProviderID {
				items[i].OrderID = createdOrder.ID
				break
			}
		}

		store[v.ProductOptionID] += v.Quantity
	}
	ok, invalidKey := quantityStore.Reduce(store)

	for invalidKey != 0 {
		var option entities2.ProductOption
		err = query.Table(entities2.ProductOption{}.TableName()).
			Where("id = ?", invalidKey).First(&option).Error
		if err != nil {
			query.Rollback()
			return nil, err
		}

		quantityStore.Add(invalidKey, option.Quantity)
		ok, invalidKey = quantityStore.Reduce(store)
	}

	if !ok {
		query.Rollback()
		return nil, fmt.Errorf("product is not have enough quantity")
	}

	for _, v := range items {
		err = query.Table(entities2.ProductOption{}.TableName()).
			Where("id = ?", v.ProductOptionID).
			UpdateColumn("quantity", gorm.Expr("quantity - ?", v.Quantity)).
			Error
		if err != nil {
			query.Rollback()
			quantityStore.Restore(store)
			return nil, err
		}
	}

	err = query.Table(entities.OrderItem{}.TableName()).Create(&items).Error
	if err != nil {
		query.Rollback()
		quantityStore.Restore(store)
		return nil, err
	}

	query.Commit()

	return orders, nil
}
