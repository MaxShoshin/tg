// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seniorGolang/tg/v2/example/interfaces"
)

type httpExampleRPC struct {
	errorHandler     ErrorHandler
	maxBatchSize     int
	maxParallelBatch int
	svc              *serverExampleRPC
	base             interfaces.ExampleRPC
}

func NewExampleRPC(svcExampleRPC interfaces.ExampleRPC) (srv *httpExampleRPC) {

	srv = &httpExampleRPC{
		base: svcExampleRPC,
		svc:  newServerExampleRPC(svcExampleRPC),
	}
	return
}

func (http *httpExampleRPC) Service() MiddlewareSetExampleRPC {
	return http.svc
}

func (http *httpExampleRPC) WithLog() *httpExampleRPC {
	http.svc.WithLog()
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
	route.Post("/api/v1/exampleRPC", http.serveBatch)
	route.Post("/api/v1/exampleRPC/test", http.serveTest)
	route.Post("/api/v1/exampleRPC/test2", http.serveTest2)
}
