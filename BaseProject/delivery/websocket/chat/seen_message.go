package chat

import (
	"context"
	"encoding/json"
	"fmt"
	socketIO "github.com/eNViDAT0001/Thesis/Backend/delivery/websocket/chat/io"
	"github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"strconv"
)

func (s *socketChatHandler) SeenMessage() io.EventHandler {
	return func(event *io.Event, c io.Client) (outGoingEvent io.Event, targetID string, err error) {
		var connectPayload socketIO.SeenMessageEvent
		if err := json.Unmarshal(event.Payload, &connectPayload); err != nil {
			return outGoingEvent, targetID, fmt.Errorf("bad payload in request: %v", err)
		}

		err = s.chatUC.SeenMessages(context.Background(), connectPayload.ID, connectPayload.UserID, connectPayload.ToID)
		if err != nil {
			return outGoingEvent, targetID, err
		}

		data, err := json.Marshal(connectPayload)
		if err != nil {
			return outGoingEvent, targetID, fmt.Errorf("failed to marshal broadcast message: %v", err)
		}
		outGoingEvent.Payload = data
		outGoingEvent.Type = io.ChatSeenMessage

		return outGoingEvent, strconv.Itoa(int(connectPayload.ToID)), nil
	}
}
