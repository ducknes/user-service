package usercontext

import (
	"context"
	"github.com/GOAT-prod/goatlogger"
	"user-service/settings"
)

type UserContext struct {
	ctx    context.Context
	logger *goatlogger.Logger
}

func New() UserContext {
	logger := goatlogger.New(settings.AppName())

	return UserContext{
		ctx:    context.Background(),
		logger: &logger,
	}
}

func (uc *UserContext) SetCtx(ctx context.Context) {
	uc.ctx = ctx
}

func (uc *UserContext) Ctx() context.Context {
	return uc.ctx
}

func (uc *UserContext) Log() *goatlogger.Logger {
	return uc.logger
}

func (uc *UserContext) SetLogTag(tag string) {
	uc.logger.SetTag(tag)
}
