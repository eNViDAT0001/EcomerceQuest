package smtp

import "context"

type UseCase interface {
	SendEmail(ctx context.Context, email string, body string) error
}
