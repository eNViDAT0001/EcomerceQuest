package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	chat2 "github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
	"gorm.io/gorm"
)

type chatUseCase struct {
	chatSto chat2.Storage
}

func (u *chatUseCase) SeenMessages(ctx context.Context, id uint, userID uint, toID uint) error {
	return u.chatSto.SeenMessages(ctx, id, userID, toID)
}

func (u *chatUseCase) Create(ctx context.Context, input io.MessageInput) (io.MessageInput, error) {
	return u.chatSto.Create(ctx, input)
}

func (u *chatUseCase) Update(ctx context.Context, id uint, input io.MessageUpdateInput) error {
	return u.chatSto.Update(ctx, id, input)
}

func (u *chatUseCase) Delete(ctx context.Context, id uint, userID uint) error {
	return u.chatSto.Delete(ctx, id, userID)
}
func (u *chatUseCase) ListChannel(ctx context.Context, userID uint, filter paging.ParamsInput) (messages []io.MessageChannel, total int64, err error) {
	total, err = u.chatSto.CountListChannel(ctx, userID, filter)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	messages, err = u.chatSto.ListChannel(ctx, userID, filter)
	if err != nil {
		return nil, 0, err
	}

	return messages, total, err
}
func (u *chatUseCase) List(ctx context.Context, input io.ListMessageInput) (messages []entities.Message, total int64, err error) {
	total, err = u.chatSto.CountList(ctx, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	messages, err = u.chatSto.List(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return messages, total, err
}

func NewChatUseCase(chatSto chat2.Storage) chat2.UseCase {
	return &chatUseCase{chatSto: chatSto}
}
