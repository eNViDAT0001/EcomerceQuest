package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/chat/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat/storage/io"
	"github.com/jinzhu/copier"
)

func UpdateMessageReqToUpdateMessageForm(input *ioHttpHandler.MessageUpdateReq) (ioSto.MessageUpdateInput, error) {
	var result ioSto.MessageUpdateInput
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
func CreateMessageReqToCreateMessageForm(input *ioHttpHandler.MessageCreateReq) (ioSto.MessageInput, error) {
	var result ioSto.MessageInput
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
