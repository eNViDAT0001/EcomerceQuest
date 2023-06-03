package request_contact

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
)

type UseCase interface {
	CreateRequest(ctx context.Context, input io.CreateRequestContact) (id uint, err error)
	ListRequest(ctx context.Context, filter paging.ParamsInput) (requests []entities.RequestContact, total int64, err error)
	DeleteRequest(ctx context.Context, id uint) (err error)
	UpdateRequestContact(ctx context.Context, id uint, input io.UpdateRequestContactForm) (err error)
}
