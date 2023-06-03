package smtp

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	ioRequest "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact/storage/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
)

type UseCase interface {
	SendEmail(ctx context.Context, input ioSto.EmailForm) error
	CreateEmail(ctx context.Context, email ioSto.CreateEmail, request ioRequest.CreateRequestContact) (id uint, err error)
	ListEmails(ctx context.Context, filter paging.ParamsInput) (emails []entities.Email, total int64, err error)

	ResendEmails(ctx context.Context, status string) (err error)
}
