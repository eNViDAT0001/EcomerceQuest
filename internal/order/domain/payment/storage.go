package payment

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/payment/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

type Storage interface {
	CreatePayment(ctx context.Context, input io.CreatePaymentForm) (io.CreatePaymentForm, error)
	GetPaymentByID(ctx context.Context, paymentID string) (entities.Payment, error)
	GetPaymentByOrderID(ctx context.Context, paymentID string) (entities.Payment, error)
}
