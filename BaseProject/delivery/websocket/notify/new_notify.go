package notify

import (
	"context"
	"encoding/json"
	"fmt"
	notifySto "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"strconv"
)

func (s socketNotificationHandler) NewNotification() io.EventHandler {
	return func(event *io.Event, c io.Client) (outGoingEvent io.Event, targetID string, err error) {
		var connectPayload notifySto.NotificationInput
		if err = json.Unmarshal(event.Payload, &connectPayload); err != nil {
			return outGoingEvent, targetID, fmt.Errorf("bad payload in request: %v", err)
		}

		message, err := s.notifyUC.CreateNotification(context.Background(), connectPayload)
		if err != nil {
			return outGoingEvent, targetID, err
		}

		data, err := json.Marshal(message)
		if err != nil {
			return outGoingEvent, targetID, fmt.Errorf("failed to marshal broadcast message: %v", err)
		}
		outGoingEvent.Payload = data
		outGoingEvent.Type = io.NotificationNew

		return outGoingEvent, strconv.Itoa(int(connectPayload.UserID)), nil
	}
}
