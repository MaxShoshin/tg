// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"

	"github.com/seniorGolang/tg/v2/example/interfaces"
)

type serverExampleRPC struct {
	svc   interfaces.ExampleRPC
	test  ExampleRPCTest
	test2 ExampleRPCTest2
}

type MiddlewareSetExampleRPC interface {
	Wrap(m MiddlewareExampleRPC)
	WrapTest(m MiddlewareExampleRPCTest)
	WrapTest2(m MiddlewareExampleRPCTest2)

	WithTrace()
	WithMetrics()
	WithLog()
}

func newServerExampleRPC(svc interfaces.ExampleRPC) *serverExampleRPC {
	return &serverExampleRPC{
		svc:   svc,
		test:  svc.Test,
		test2: svc.Test2,
	}
}

func (srv *serverExampleRPC) Wrap(m MiddlewareExampleRPC) {
	srv.svc = m(srv.svc)
	srv.test = srv.svc.Test
	srv.test2 = srv.svc.Test2
}

func (srv *serverExampleRPC) Test(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (ret1 int, ret2 string, err error) {
	return srv.test(ctx, arg0, arg1, opts...)
}

func (srv *serverExampleRPC) Test2(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (ret1 int, ret2 string, err error) {
	return srv.test2(ctx, arg0, arg1, opts...)
}

func (srv *serverExampleRPC) WrapTest(m MiddlewareExampleRPCTest) {
	srv.test = m(srv.test)
}

func (srv *serverExampleRPC) WrapTest2(m MiddlewareExampleRPCTest2) {
	srv.test2 = m(srv.test2)
}

func (srv *serverExampleRPC) WithTrace() {
	srv.Wrap(traceMiddlewareExampleRPC)
}

func (srv *serverExampleRPC) WithMetrics() {
	srv.Wrap(metricsMiddlewareExampleRPC)
}

func (srv *serverExampleRPC) WithLog() {
	srv.Wrap(loggerMiddlewareExampleRPC())
}
