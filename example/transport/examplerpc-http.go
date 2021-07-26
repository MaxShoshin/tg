// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"github.com/seniorGolang/tg/v2/example/interfaces"
)

type httpExampleRPC struct {
	log          zerolog.Logger
	errorHandler ErrorHandler
	svc          *serverExampleRPC
	base         interfaces.ExampleRPC
}

func NewExampleRPC(log zerolog.Logger, svcExampleRPC interfaces.ExampleRPC) (srv *httpExampleRPC) {

	srv = &httpExampleRPC{
		base: svcExampleRPC,
		log:  log,
		svc:  newServerExampleRPC(svcExampleRPC),
	}
	return
}

func (http httpExampleRPC) Service() MiddlewareSetExampleRPC {
	return http.svc
}

func (http *httpExampleRPC) WithLog(log zerolog.Logger) *httpExampleRPC {
	http.svc.WithLog(log)
	return http
}

func (http *httpExampleRPC) WithTrace() *httpExampleRPC {
	http.svc.WithTrace()
	return http
}

func (http *httpExampleRPC) WithMetrics() *httpExampleRPC {
	http.svc.WithMetrics()
	return http
}

func (http *httpExampleRPC) WithErrorHandler(handler ErrorHandler) *httpExampleRPC {
	http.errorHandler = handler
	return http
}

func (http *httpExampleRPC) SetRoutes(route *fiber.App) {
	route.Post("/examplerpc", http.serveBatch)
	route.Post("/exampleRPC/test", http.serveTest)
}
