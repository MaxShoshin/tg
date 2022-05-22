// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/seniorGolang/dumper/viewer"
	"github.com/seniorGolang/tg/v2/example/interfaces"
	"github.com/seniorGolang/tg/v2/example/interfaces/types"
	"time"
)

type loggerUser struct {
	next interfaces.User
}

func loggerMiddlewareUser() MiddlewareUser {
	return func(next interfaces.User) interfaces.User {
		return &loggerUser{next: next}
	}
}

func (m loggerUser) GetUser(ctx context.Context, cookie string, userAgent string) (user *types.User, err error) {
	logger := log.Ctx(ctx).With().Str("service", "User").Str("method", "getUser").Logger()
	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"request": viewer.Sprintf("%+v", requestUserGetUser{
				Cookie:    cookie,
				UserAgent: userAgent,
			}),
			"response": viewer.Sprintf("%+v", responseUserGetUser{User: user}),
			"took":     time.Since(begin).String(),
		}
		if err != nil {
			logger.Error().Err(err).Fields(fields).Msg("call getUser")
			return
		}
		logger.Info().Fields(fields).Msg("call getUser")
	}(time.Now())
	return m.next.GetUser(ctx, cookie, userAgent)
}

func (m loggerUser) UploadFile(ctx context.Context, fileBytes []byte) (err error) {
	logger := log.Ctx(ctx).With().Str("service", "User").Str("method", "uploadFile").Logger()
	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"request":  viewer.Sprintf("%+v", requestUserUploadFile{FileBytes: fileBytes}),
			"response": viewer.Sprintf("%+v", responseUserUploadFile{}),
			"took":     time.Since(begin).String(),
		}
		if err != nil {
			logger.Error().Err(err).Fields(fields).Msg("call uploadFile")
			return
		}
		logger.Info().Fields(fields).Msg("call uploadFile")
	}(time.Now())
	return m.next.UploadFile(ctx, fileBytes)
}

func (m loggerUser) CustomResponse(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (err error) {
	logger := log.Ctx(ctx).With().Str("service", "User").Str("method", "customResponse").Logger()
	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"request": viewer.Sprintf("%+v", requestUserCustomResponse{
				Arg0: arg0,
				Arg1: arg1,
				Opts: opts,
			}),
			"response": viewer.Sprintf("%+v", responseUserCustomResponse{}),
			"took":     time.Since(begin).String(),
		}
		if err != nil {
			logger.Error().Err(err).Fields(fields).Msg("call customResponse")
			return
		}
		logger.Info().Fields(fields).Msg("call customResponse")
	}(time.Now())
	return m.next.CustomResponse(ctx, arg0, arg1, opts...)
}

func (m loggerUser) CustomHandler(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (err error) {
	logger := log.Ctx(ctx).With().Str("service", "User").Str("method", "customHandler").Logger()
	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"request": viewer.Sprintf("%+v", requestUserCustomHandler{
				Arg0: arg0,
				Arg1: arg1,
				Opts: opts,
			}),
			"response": viewer.Sprintf("%+v", responseUserCustomHandler{}),
			"took":     time.Since(begin).String(),
		}
		if err != nil {
			logger.Error().Err(err).Fields(fields).Msg("call customHandler")
			return
		}
		logger.Info().Fields(fields).Msg("call customHandler")
	}(time.Now())
	return m.next.CustomHandler(ctx, arg0, arg1, opts...)
}
