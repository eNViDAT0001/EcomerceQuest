package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
)

type requestContactUseCase struct {
	requestContactSto request_contact.Storage
	notifyUC          notification.UseCase
}

func (u *requestContactUseCase) UpdateRequestContact(ctx context.Context, id uint, input io.UpdateRequestContactForm) (err error) {
	return u.requestContactSto.UpdateRequestContact(ctx, id, input)
}

func (u *requestContactUseCase) CreateRequest(ctx context.Context, input io.CreateRequestContact) (id uint, err error) {
	id, err = u.requestContactSto.CreateRequest(ctx, input)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *requestContactUseCase) ListRequest(ctx context.Context, filter paging.ParamsInput) (requests []entities.RequestContact, total int64, err error) {
	return u.requestContactSto.ListRequest(ctx, filter)
}

func (u *requestContactUseCase) DeleteRequest(ctx context.Context, id uint) (err error) {
	return u.requestContactSto.DeleteRequest(ctx, id)
}

func NewRequestContactUseCase(requestContactSto request_contact.Storage, notifyUC notification.UseCase) request_contact.UseCase {
	return &requestContactUseCase{requestContactSto: requestContactSto, notifyUC: notifyUC}
}
