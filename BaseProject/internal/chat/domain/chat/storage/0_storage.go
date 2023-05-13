package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
)

type chatStorage struct {
}

func (s *chatStorage) ListChannel(ctx context.Context, userID uint, filter paging.ParamsInput) ([]io.MessageChannel, error) {
	result := make([]io.MessageChannel, 0)
	db := wrap_gorm.GetDB()

	channels := make([]io.Channel, 0)
	err := db.Raw("SELECT DISTINCT Message.user_id, Message.to_user_id "+
		"FROM Message WHERE (Message.user_id, Message.to_user_id) "+
		"IN ( SELECT Message.user_id, Message.to_user_id FROM Message GROUP BY Message.user_id, Message.to_user_id ORDER BY Message.created_at DESC, Message.seen DESC ) "+
		"AND (Message.user_id = ? OR Message.to_user_id = ?) LIMIT ? OFFSET ?", userID, userID, filter.PerPage(), paging.Offset(filter.Current(), filter.PerPage())).
		Find(&channels).
		Error
	if err != nil {
		return nil, err
	}

	senderChannels := map[uint][]uint{}
	receiverChannels := map[uint][]uint{}
	for _, channel := range channels {
		flag := false
		for _, v := range senderChannels[channel.UserID] {
			if v == channel.UserID {
				receiverChannels[channel.UserID] = append(receiverChannels[channel.UserID], v)
				flag = true
				break
			}
		}
		if !flag {
			senderChannels[channel.UserID] = append(senderChannels[channel.UserID], channel.ToUserID)
		}
	}

	for k, v := range senderChannels {
		message := io.MessageChannel{}
		query := db.Model(entities.Message{}).
			Select("User.id AS channel_id, User.name, User.avatar, Message.id, Message.user_id, Message.content, Message.to_user_id, Message.seen, Message.type, Message.created_at").
			Where("Message.user_id = ? AND Message.to_user_id IN ?", k, v)
		if fromIDs, ok := receiverChannels[k]; ok {
			query = query.Or("Message.user_id IN ? AND Message.to_user_id = ?", fromIDs, k).
				Joins("JOIN User ON User.id = Message.to_user_id")
		} else {
			query = query.Joins("JOIN User ON User.id = Message.user_id")
		}
		query = query.Order("Message.id DESC, Message.seen DESC").Scan(&message)
		if query.Error != nil {
			return nil, query.Error
		}
		result = append(result, message)
	}
	return result, nil
}
func remove(slice []uint, s uint) []uint {
	index := 0
	for i, u := range slice {
		if u == s {
			index = i
			break
		}
	}

	if index == 0 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
func (s *chatStorage) ListMessageByChannel(ctx context.Context, userID uint, toID uint, filter *paging.ParamsInput) ([]entities.Message, error) {
	result := make([]entities.Message, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Message{}).
		Where("user_id = ? AND to_user_id = ?", userID, toID)

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
		Select("COUNT(DISTINCT Message.user_id, Message.to_user_id)").
		Where("user_id = ? OR to_user_id = ?", userID, userID)

	paging_query.SetCountListPagingQuery(&filter, entities.Message{}.TableName(), query)

	err := query.Order("created_at DESC, seen DESC").Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *chatStorage) Create(ctx context.Context, input io.MessageInput) (io.MessageInput, error) {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Message{}.TableName()).Create(&input).Error
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
		Where("user_id = ?", userID).
		Where("to_user_id = ?", userID).
		Where("seen = ?", false).
		Update("seen", true).Error
	return err
}

func (s chatStorage) Delete(ctx context.Context, id uint, userID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Message{}.TableName()).
		Where("id = ?", id).
		Where("user_id = ?", userID).
		Delete(entities.Message{}).Error
	return err
}

func (s chatStorage) List(ctx context.Context, input io.ListMessageInput) ([]entities.Message, error) {
	result := make([]entities.Message, 0)
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Message{}).Where("(user_id = ? AND to_user_id = ?) OR (user_id = ? AND to_user_id = ?)", input.UserIDA, input.UserIDB, input.UserIDB, input.UserIDA)

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

	query := db.Model(entities.Message{}).Where("(user_id = ? AND to_user_id = ?) OR (user_id = ? AND to_user_id = ?)", input.UserIDA, input.UserIDB, input.UserIDB, input.UserIDA)

	paging_query.SetCountListPagingQuery(&input.Paging, entities.Message{}.TableName(), query)

	err := query.Order("created_at DESC, seen DESC").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func NewChatStorage() chat.Storage {
	return &chatStorage{}
}
