package service

import (
	"context"
	"user_svc/pkg/model"

	log "github.com/go-kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserSvcService) UserSvcService

type loggingMiddleware struct {
	logger log.Logger
	next   UserSvcService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserSvcService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserSvcService) UserSvcService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, email string, password string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Create", "email", email, "password", password, "rs", rs, "err", err)
	}()
	return l.next.Create(ctx, email, password)
}

func (l loggingMiddleware) Get(ctx context.Context) (rs []model.User, err error) {
	defer func() {
		l.logger.Log("method", "Get", "rs", rs, "err", err)
	}()
	return l.next.Get(ctx)
}
