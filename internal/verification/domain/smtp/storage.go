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
}
