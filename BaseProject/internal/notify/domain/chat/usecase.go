package chat

import (
	"context"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type UseCase interface {
	Create(ctx context.Context, input io2.MessageInput) error
	Update(ctx context.Context, id uint, input io2.MessageUpdateInput) error
	Delete(ctx context.Context, id uint, userID uint) error
	SeenMessages(ctx context.Context, id uint) error
	List(ctx context.Context, input io2.ListMessageInput) (messages []entities.Message, total int64, err error)
}
