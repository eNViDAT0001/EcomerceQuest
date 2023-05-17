package chat

import (
	context "context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type notificationHandler struct {
	notifyUC notification.UseCase
}

func (s *notificationHandler) SeenNotification() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		id, err := strconv.Atoi(cc.Param("notify_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		err = s.notifyUC.SeenNotification(newCtx, uint(id), uint(userID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Message Success")
	}
}
func (s *notificationHandler) SeenAllNotification() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		err = s.notifyUC.SeenAllNotification(newCtx, uint(userID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Message Success")
	}
}

func (s *notificationHandler) ListNotifications() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Notification{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		inputRepo := io.ListNotifyInput{
			UserID: uint(userID),
			Paging: &paginator,
		}

		result, total, err := s.notifyUC.ListNotification(newCtx, inputRepo)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent(err)
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(result) > 0 {
			paginator.Marker = int(result[len(result)-1].ID)
		}
		cc.OkPaging(paginator, result)
	}
}
func (s *notificationHandler) ListNotificationFullView() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Notification{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		inputRepo := io.ListNotifyInput{
			UserID: uint(userID),
			Paging: &paginator,
		}

		result, total, totalUnseens, err := s.notifyUC.ListNotificationFullView(newCtx, inputRepo)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent(err)
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(result) > 0 {
			paginator.Marker = int(result[len(result)-1].ID)
		}
		extraData := request.ExtraData{
			Name: "total_unseen",
			Data: totalUnseens,
		}
		cc.OkPaging(paginator, result, extraData)
	}
}

func NewNotificationHandler(notifyUC notification.UseCase) notification.HttpHandler {
	return &notificationHandler{notifyUC: notifyUC}
}
