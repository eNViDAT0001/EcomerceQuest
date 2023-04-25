package smtp

import (
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/event_background"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/usecase/io"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (s *smtpHandler) CreateResetPassCode() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		email := cc.Query("email")

		token, code, err := s.jwtUC.GenerateSmtpCode(newCtx, email)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result := map[string]interface{}{
			"token": token,
		}
		cc.Ok(result)

		mail := io.EmailForm{
			Subject: "Reset password",
			Content: fmt.Sprintf("<h1>%s</h1>", code),
			To:      []string{email},
		}

		job := event_background.NewJob(func(ctx context.Context) error {
			return s.smtpUC.SendEmail(ctx, mail)
		})
		event_background.GetBackGroundJobs().Group <- event_background.NewJobs(true, job)

	}
}
