package chat

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/chat/convert"
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/chat/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_params"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
	"github.com/eNViDAT0001/Thesis/Backend/socket"
	io2 "github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type chatHandler struct {
	chatUC chat.UseCase
}

func (s *chatHandler) SeenMessages() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		id, err := strconv.Atoi(cc.Param("message_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		toID, err := strconv.Atoi(cc.Param("to_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		err = s.chatUC.SeenMessages(newCtx, uint(id), uint(userID), uint(toID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Message Success")
	}
}

func (s *chatHandler) CreateMessage() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input ioHandler.MessageCreateReq
		if err := cc.ShouldBind(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputRepo, err := convert.CreateMessageReqToCreateMessageForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		message, err := s.chatUC.Create(newCtx, inputRepo)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = socket.GetManager().
			EmitNotify(io2.ChatNewMessage, message, strconv.Itoa(int(message.ToUserID)))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(message)
	}
}

func (s *chatHandler) UpdateMessages() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		id, err := strconv.Atoi(cc.Param("id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		var input ioHandler.MessageUpdateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputRepo, err := convert.UpdateMessageReqToUpdateMessageForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		err = s.chatUC.Update(newCtx, uint(id), inputRepo)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Message Success")
	}
}

func (s *chatHandler) ListMessages() func(ctx *gin.Context) {
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

		inValidField, val := paging_params.ValidateFilter(paginator.Filter, entities.Message{})
		if len(inValidField) > 0 {
			cc.ResponseError(request.NewBadRequestError(inValidField, val, "invalid key and value"))
			return
		}

		userA, err := strconv.Atoi(cc.Param("user_id_a"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		userB, err := strconv.Atoi(cc.Param("user_id_b"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		inputRepo := io.ListMessageInput{
			UserIDA: userA,
			UserIDB: userB,
			Paging:  paginator,
		}

		result, total, err := s.chatUC.List(newCtx, inputRepo)
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
func (s *chatHandler) ListChannel() func(ctx *gin.Context) {
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

		inValidField, val := paging_params.ValidateFilter(paginator.Filter, entities.Message{})
		if len(inValidField) > 0 {
			cc.ResponseError(request.NewBadRequestError(inValidField, val, "invalid key and value"))
			return
		}

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, total, err := s.chatUC.ListChannel(newCtx, uint(userID), paginator)
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
func NewChatHandler(chatUC chat.UseCase) chat.HttpHandler {
	return &chatHandler{chatUC: chatUC}
}
