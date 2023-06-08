package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/payment"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/payment/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

type paymentUseCase struct {
	paymentSto payment.Storage
}

func (u *paymentUseCase) CreatePayment(ctx context.Context, input io.CreatePaymentForm) (io.CreatePaymentForm, error) {
	return u.paymentSto.CreatePayment(ctx, input)
}

func (u *paymentUseCase) GetPaymentByID(ctx context.Context, paymentID string) (entities.Payment, error) {
	return u.paymentSto.GetPaymentByID(ctx, paymentID)
}
func (u *paymentUseCase) GetPaymentByOrderID(ctx context.Context, orderID string) (entities.Payment, error) {
	return u.paymentSto.GetPaymentByOrderID(ctx, orderID)
}

func NewPaymentUseCase(paymentSto payment.Storage) payment.UseCase {
	return &paymentUseCase{paymentSto: paymentSto}
}
