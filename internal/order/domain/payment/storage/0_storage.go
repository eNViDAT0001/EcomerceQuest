package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/payment"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/payment/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

type paymentStorage struct {
}

func (s *paymentStorage) GetPaymentByID(ctx context.Context, paymentID string) (entities.Payment, error) {
	result := entities.Payment{}
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Payment{}).Where("id = ?", paymentID).First(&result).Error

	return result, err
}
func (s *paymentStorage) GetPaymentByOrderID(ctx context.Context, orderID string) (entities.Payment, error) {
	result := entities.Payment{}
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Payment{}).
		Joins("`Order` ON `Order`.payment_id = Payment.id").
		Where("`Order`.`id` = ?", orderID).
		First(&result).Error

	return result, err
}

func (s *paymentStorage) CreatePayment(ctx context.Context, input io.CreatePaymentForm) (io.CreatePaymentForm, error) {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Payment{}.TableName()).Create(&input).Error

	return input, err
}

func (s *paymentStorage) UpdatePayment(ctx context.Context, paymentID string, form io.UpdatePaymentForm) error {
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Payment{}).Where("id = ?", paymentID).Updates(&form).Error

	return err
}

func NewPaymentStorage() payment.Storage {
	return &paymentStorage{}
}
