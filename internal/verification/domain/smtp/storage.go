package smtp

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
)

type Storage interface {
	CreateEmail(ctx context.Context, email io.CreateEmail) (id uint, err error)
	CountListEmail(ctx context.Context, filter paging.ParamsInput) (int64, error)
	ListEmail(ctx context.Context, filter paging.ParamsInput) ([]entities.Email, error)

	StoreEmail(ctx context.Context, email io.EmailForm) (io.EmailForm, error)
	UpdateSentByID(ctx context.Context, id uint, status string) (err error)
	GetEmailsByStatus(ctx context.Context, status string) (emails []entities.Smtp, err error)
	DeleteUnsentEmails(ctx context.Context) (err error)
}
