package notify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	socketIO "github.com/eNViDAT0001/Thesis/Backend/delivery/websocket/notify/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_params"
	notifySto "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	notifyEntities "github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
	"github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"strconv"
)

func (s socketNotificationHandler) UnseenNotification() io.EventHandler {
	return func(event *io.Event, c io.Client) (outGoingEvent io.Event, targetID string, err error) {
		var connectPayload socketIO.NotifyUnseenEvent
		if err := json.Unmarshal(event.Payload, &connectPayload); err != nil {
			return outGoingEvent, targetID, fmt.Errorf("bad payload in request: %v", err)
		}

		fields := connectPayload.Fields
		search := connectPayload.Search
		sort := connectPayload.Sort
		filter := paging_params.NewFilterBuilder().
			WithSorts(sort).
			WithFields(fields).
			WithSearch(search).
			Build()
		inValidField, val := paging_params.ValidateFilter(filter, notifyEntities.Notification{})
		if len(inValidField) > 0 {
			return outGoingEvent, targetID, errors.New(inValidField + ": " + val + "-> invalid key and value")
		}

		if connectPayload.Type != paging.CursorPaging {
			connectPayload.Type = paging.NUM_PAGING
		}
		if connectPayload.Limit == 0 {
			connectPayload.Limit = 20
		}
		inputUC := notifySto.ListNotifyInput{
			UserID: connectPayload.UserID,
			Paging: paging.ParamsInput{
				Marker: connectPayload.Marker,
				Limit:  connectPayload.Limit,
				Total:  connectPayload.Total,
				Type:   connectPayload.Type,
				Filter: filter,
			},
		}
		notifications, total, err := s.notifyUC.ListNotification(context.Background(), inputUC)
		if err != nil {
			return outGoingEvent, targetID, err
		}
		inputUC.Paging.Total = int(total)
		if inputUC.Paging.Type == paging.CursorPaging && len(notifications) > 0 {
			inputUC.Paging.Marker = int(notifications[len(notifications)-1].ID)
		}
		newPaginator := paging.NewPaginator(inputUC.Paging)
		// Prepare an Outgoing Message to others
		var broadMessage socketIO.NewNotifyUnseenEvent
		broadMessage.Notifications = notifications
		broadMessage.Paging = newPaginator

		data, err := json.Marshal(broadMessage)
		if err != nil {
			return outGoingEvent, targetID, fmt.Errorf("failed to marshal broadcast message: %v", err)
		}
		outGoingEvent.Payload = data
		outGoingEvent.Type = io.NotifyUnseenMessage

		return outGoingEvent, strconv.Itoa(int(connectPayload.UserID)), nil
	}
}
