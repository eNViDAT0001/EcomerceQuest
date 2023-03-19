package smtp

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *smtpHandler) GetReportEmail() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		//newCtx := context.Background()

		cc.NotFound(request.NewNotFoundError("", "", "Not implemented"))
	}
}
