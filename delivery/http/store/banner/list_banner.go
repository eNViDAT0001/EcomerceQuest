package banner

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	paging_query "github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s bannerHandler) ListBanner() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Banner{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		comments, total, err := s.bannerUC.ListBanner(newCtx, paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(comments) > 0 {
			paginator.Marker = int(comments[len(comments)-1].ID)
		}

		cc.OkPaging(paginator, comments)
	}
}
