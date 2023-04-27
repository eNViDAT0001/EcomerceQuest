package chat

import (
	"context"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type Storage interface {
	Create(ctx context.Context, input io2.MessageInput) error
	Update(ctx context.Context, id uint, input io2.MessageUpdateInput) error
	SeenMessages(ctx context.Context, id uint) error
	Delete(ctx context.Context, id uint, userID uint) error
	List(ctx context.Context, input io2.ListMessageInput) ([]entities.Message, error)
	CountList(ctx context.Context, input io2.ListMessageInput) (int64, error)
}
