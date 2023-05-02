package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
	"gorm.io/gorm"
)

type chatUseCase struct {
	chatSto chat.Storage
}

func (u *chatUseCase) SeenMessages(ctx context.Context, id uint, userID uint, toID uint) error {
	return u.chatSto.SeenMessages(ctx, id, userID, toID)
}

func (u *chatUseCase) Create(ctx context.Context, input io2.MessageInput) error {
	return u.chatSto.Create(ctx, input)
}

func (u *chatUseCase) Update(ctx context.Context, id uint, input io2.MessageUpdateInput) error {
	return u.chatSto.Update(ctx, id, input)
}

func (u *chatUseCase) Delete(ctx context.Context, id uint, userID uint) error {
	return u.chatSto.Delete(ctx, id, userID)
}

func (u *chatUseCase) List(ctx context.Context, input io2.ListMessageInput) (messages []entities.Message, total int64, err error) {
	total, err = u.chatSto.CountList(ctx, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	products, err := u.chatSto.List(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}

func NewChatUseCase(chatSto chat.Storage) chat.UseCase {
	return &chatUseCase{chatSto: chatSto}
}
