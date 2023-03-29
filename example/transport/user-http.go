// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/gofiber/fiber/v2"
	implement "github.com/seniorGolang/tg/v2/example/implement"
	"github.com/seniorGolang/tg/v2/example/interfaces"
)

type httpUser struct {
	errorHandler     ErrorHandler
	maxBatchSize     int
	maxParallelBatch int
	svc              *serverUser
	base             interfaces.User
}

func NewUser(svcUser interfaces.User) (srv *httpUser) {

	srv = &httpUser{
		base: svcUser,
		svc:  newServerUser(svcUser),
	}
	return
}

func (http *httpUser) Service() MiddlewareSetUser {
	return http.svc
}

func (http *httpUser) WithLog() *httpUser {
	http.svc.WithLog()
	return http
}

func (http *httpUser) WithTrace() *httpUser {
	http.svc.WithTrace()
	return http
}

func (http *httpUser) WithMetrics() *httpUser {
	http.svc.WithMetrics()
	return http
}

func (http *httpUser) WithErrorHandler(handler ErrorHandler) *httpUser {
	http.errorHandler = handler
	return http
}

func (http *httpUser) SetRoutes(route *fiber.App) {
	route.Get("/api/v1/api/v2/user/info", http.serveGetUser)
	route.Patch("/api/v1/api/v2/user/custom/response", http.serveCustomResponse)
	route.Delete("/api/v1/api/v2/user/custom", func(ctx *fiber.Ctx) (err error) {
		return implement.CustomHandler(ctx, http.base)
	})
}
