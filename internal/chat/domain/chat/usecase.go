package chat

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
)

type UseCase interface {
	Create(ctx context.Context, input io.MessageInput) (io.MessageInput, error)
	Update(ctx context.Context, id uint, input io.MessageUpdateInput) error
	Delete(ctx context.Context, id uint, userID uint) error
	SeenMessages(ctx context.Context, id uint, userID uint, toID uint) error
	List(ctx context.Context, input io.ListMessageInput) (messages []entities.Message, total int64, err error)
	ListChannel(ctx context.Context, userID uint, filter paging.ParamsInput) (chatRooms []io.ChatRoom, total int64, err error)
	GetByID(ctx context.Context, fromUserID uint, toUserID uint) (entities.ChatRoom, error)
}
