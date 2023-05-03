package chat

import (
	context "context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_params"
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
	//TODO implement me
	panic("implement me")
}

func (s *notificationHandler) ListNotifications() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator := paging.ParamsInput{}
		if err := cc.BindQuery(&paginator); err != nil {
			cc.BadRequest(err)
			return
		}

		search := cc.QueryArray("search[]")
		fields := cc.QueryArray("fields[]")
		sort := cc.QueryArray("sorts[]")

		paginator.Filter = paging_params.NewFilterBuilder().
			WithSearch(search).
			WithFields(fields).
			WithSorts(sort).
			Build()

		inValidField, val := paging_params.ValidateFilter(paginator.Filter, entities.Notification{})
		if len(inValidField) > 0 {
			cc.ResponseError(request.NewBadRequestError(inValidField, val, "invalid key and value"))
			return
		}

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		inputRepo := io.ListNotifyInput{
			UserID: uint(userID),
			Paging: paging.ParamsInput{},
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

func NewNotificationHandler(notifyUC notification.UseCase) notification.HttpHandler {
	return &notificationHandler{notifyUC: notifyUC}
}
