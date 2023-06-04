package smtp

import (
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/cache"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
	"time"
)

func UseToken(ctx context.Context, token string) bool {
	_, err := cache.GetRedis().Get(ctx, "token_"+token)
	if err == redis.Nil {
		err = cache.GetRedis().Set(ctx, "token_"+token, true, 5*time.Minute)
		if err != nil {
			log.Printf("Error when set token to redis: %v", err)
		}

		return true
	}

	return false
}
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

		mail := io.EmailForm{
			Subject: "Reset password",
			Content: fmt.Sprintf("<h1>%s</h1>", code),
			To:      []string{email},
		}

		err = s.smtpUC.SendEmail(newCtx, mail)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result := map[string]interface{}{
			"token": token,
		}
		cc.Ok(result)

	}
}
