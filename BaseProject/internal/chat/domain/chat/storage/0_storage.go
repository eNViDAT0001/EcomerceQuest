package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
)

type chatStorage struct {
	userSto user.Storage
}

func (s *chatStorage) ListChannel(ctx context.Context, userID uint, filter paging.ParamsInput) ([]io.ChatRoom, error) {
	chatRooms := make([]io.ChatRoom, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.ChatRoom{}).
		Select("ChatRoom.id, Message.from_user_id, Message.to_user_id, Message.content, Message.seen, Message.type, Message.created_at").
		Joins("JOIN Message ON Message.chat_room_id = ChatRoom.id").
		Where("Message.id = (SELECT MAX(Message.id) FROM Message WHERE Message.chat_room_id = ChatRoom.id) AND (Message.from_user_id = ? OR Message.to_user_id = ?) AND Message.deleted_at IS NULL", userID, userID)
	paging_query.SetPagingQuery(&filter, entities.Message{}.TableName(), query)

	err := query.Order("Message.id DESC").Scan(&chatRooms).Error
	if err != nil {
		return nil, err
	}
	fromUserIDsStorage := map[uint][]uint{}
	for i, chatRoom := range chatRooms {

		chatRoomExist := false
		if _, ok := fromUserIDsStorage[chatRoom.FromUserID]; ok {
			for _, toUserID := range fromUserIDsStorage[chatRoom.FromUserID] {
				if toUserID == chatRoom.ToUserID {
					chatRoomExist = true
					break
				}
			}
			if chatRoomExist {
				break
			}
		}
		if !chatRoomExist {
			fromUserIDsStorage[chatRoom.FromUserID] = append(fromUserIDsStorage[chatRoom.FromUserID], chatRoom.ToUserID)
		}

		if chatRoom.FromUserID == userID {
			channelUser, err := s.userSto.GetUserDetailByID(ctx, chatRoom.ToUserID)
			if err != nil {
				return nil, err
			}
			chatRooms[i].Name = *channelUser.Name
			chatRooms[i].Avatar = *channelUser.Avatar
			continue
		}

		channelUser, err := s.userSto.GetUserDetailByID(ctx, chatRoom.FromUserID)
		if err != nil {
			return nil, err
		}
		chatRooms[i].Name = *channelUser.Name
		chatRooms[i].Avatar = *channelUser.Avatar
	}

	return chatRooms, nil
}

func (s *chatStorage) ListMessageByChannel(ctx context.Context, userID uint, toID uint, filter *paging.ParamsInput) ([]entities.Message, error) {
	result := make([]entities.Message, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Message{}).
		Where("from_user_id = ? AND to_user_id = ?", userID, toID)

	paging_query.SetPagingQuery(filter, entities.Message{}.TableName(), query)

	err := query.Order("created_at DESC, seen DESC").Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *chatStorage) CountListChannel(ctx context.Context, userID uint, filter paging.ParamsInput) (int64, error) {
	var count int64
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Message{}).
		Select("COUNT(DISTINCT Message.from_user_id, Message.to_user_id)").
		Where("from_user_id = ? OR to_user_id = ?", userID, userID)

	paging_query.SetCountListPagingQuery(&filter, entities.Message{}.TableName(), query)

	err := query.Order("created_at DESC, seen DESC").Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *chatStorage) Create(ctx context.Context, input io.MessageInput) (io.MessageInput, error) {
	db := wrap_gorm.GetDB()
	if input.ChatRoomID != 0 {
		err := db.Table(entities.Message{}.TableName()).Create(&input).Error
		return input, err
	}

	var chatRoom entities.ChatRoom
	err := db.Table(entities.ChatRoom{}.TableName()).Create(&chatRoom).Error
	if err != nil {
		return input, err
	}

	input.ChatRoomID = chatRoom.ID
	err = db.Table(entities.Message{}.TableName()).Create(&input).Error
	return input, err
}
func (s *chatStorage) Update(ctx context.Context, id uint, input io.MessageUpdateInput) error {
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Message{}).
		Where("id = ?", id).
		Updates(&input).Error
	return err
}

func (s chatStorage) SeenMessages(ctx context.Context, id uint, userID uint, toID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Message{}).
		Where("id <= ?", id).
		Where("from_user_id = ?", userID).
		Where("to_user_id = ?", userID).
		Where("seen = ?", false).
		Update("seen", true).Error
	return err
}

func (s chatStorage) Delete(ctx context.Context, id uint, userID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Message{}.TableName()).
		Where("id = ?", id).
		Where("from_user_id = ?", userID).
		Delete(entities.Message{}).Error
	return err
}

func (s chatStorage) List(ctx context.Context, input io.ListMessageInput) ([]entities.Message, error) {
	result := make([]entities.Message, 0)
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Message{}).Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", input.UserIDA, input.UserIDB, input.UserIDB, input.UserIDA)

	paging_query.SetPagingQuery(&input.Paging, entities.Message{}.TableName(), query)

	err := query.Order("created_at DESC, seen DESC").Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s chatStorage) CountList(ctx context.Context, input io.ListMessageInput) (int64, error) {
	var count int64
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Message{}).Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", input.UserIDA, input.UserIDB, input.UserIDB, input.UserIDA)

	paging_query.SetCountListPagingQuery(&input.Paging, entities.Message{}.TableName(), query)

	err := query.Order("created_at DESC, seen DESC").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func NewChatStorage(userSto user.Storage) chat.Storage {
	return &chatStorage{userSto: userSto}
}
