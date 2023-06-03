package smtp

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type requestContactHandler struct {
	requestContactUC request_contact.UseCase
}

func (s *requestContactHandler) UpdateRequest() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		requestID, err := strconv.Atoi(cc.Param("request_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		var input io.UpdateRequestContactForm
		if err = cc.BindJSON(&input); err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.requestContactUC.UpdateRequestContact(newCtx, uint(requestID), input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete Request success")
	}
}

func (s *requestContactHandler) ListRequest() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.RequestContact{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		requests, total, err := s.requestContactUC.ListRequest(newCtx, paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(requests) > 0 {
			paginator.Marker = int(requests[len(requests)-1].ID)
		}

		cc.OkPaging(paginator, requests)
	}
}

func (s *requestContactHandler) DeleteRequest() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		requestID, err := strconv.Atoi(cc.Param("request_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		err = s.requestContactUC.DeleteRequest(newCtx, uint(requestID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete Request success")
	}
}

func NewRequestContactHandler(
	requestContactUC request_contact.UseCase,
) request_contact.HttpHandler {
	return &requestContactHandler{
		requestContactUC: requestContactUC,
	}
}
