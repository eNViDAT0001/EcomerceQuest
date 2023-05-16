package chat

import (
	"context"
	"encoding/json"
	"fmt"
	chatSto "github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"strconv"
)

func (s *socketChatHandler) NewMessage() io.EventHandler {
	return func(event *io.Event, c io.Client) (outGoingEvent io.Event, targetID string, err error) {
		var connectPayload chatSto.MessageInput
		if err := json.Unmarshal(event.Payload, &connectPayload); err != nil {
			return outGoingEvent, targetID, fmt.Errorf("bad payload in request: %v", err)
		}

		message, err := s.chatUC.Create(context.Background(), connectPayload)
		if err != nil {
			return outGoingEvent, targetID, err
		}

		data, err := json.Marshal(message)
		if err != nil {
			return outGoingEvent, targetID, fmt.Errorf("failed to marshal broadcast message: %v", err)
		}
		outGoingEvent.Payload = data
		outGoingEvent.Type = io.ChatNewMessage

		return outGoingEvent, strconv.Itoa(int(connectPayload.ToUserID)), nil
	}
}
