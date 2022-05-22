// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/gofiber/fiber/v2"
	otg "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/seniorGolang/json"
	"strings"
)

func (http *httpExampleRPC) serveTest(ctx *fiber.Ctx) (err error) {
	return http.serveMethod(ctx, "test", http.test)
}
func (http *httpExampleRPC) test(ctx context.Context, requestBase baseJsonRPC) (responseBase *baseJsonRPC) {

	var err error
	var request requestExampleRPCTest

	span := otg.SpanFromContext(ctx)
	span.SetTag("method", "test")

	if requestBase.Params != nil {
		if err = json.Unmarshal(requestBase.Params, &request); err != nil {
			ext.Error.Set(span, true)
			span.SetTag("msg", "request body could not be decoded: "+err.Error())
			return makeErrorResponseJsonRPC(requestBase.ID, parseError, "request body could not be decoded: "+err.Error(), nil)
		}
	}
	if requestBase.Version != Version {
		ext.Error.Set(span, true)
		span.SetTag("msg", "incorrect protocol version: "+requestBase.Version)
		return makeErrorResponseJsonRPC(requestBase.ID, parseError, "incorrect protocol version: "+requestBase.Version, nil)
	}

	var response responseExampleRPCTest
	response.Ret1, response.Ret2, err = http.svc.Test(ctx, request.Arg0, request.Arg1, request.Opts...)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		ext.Error.Set(span, true)
		span.SetTag("msg", err)
		span.SetTag("errData", toString(err))
		return makeErrorResponseJsonRPC(requestBase.ID, internalError, err.Error(), err)
	}
	responseBase = &baseJsonRPC{
		ID:      requestBase.ID,
		Version: Version,
	}
	if responseBase.Result, err = json.Marshal(response); err != nil {
		ext.Error.Set(span, true)
		span.SetTag("msg", "response body could not be encoded: "+err.Error())
		return makeErrorResponseJsonRPC(requestBase.ID, parseError, "response body could not be encoded: "+err.Error(), nil)
	}
	return
}
func (http *httpExampleRPC) serveMethod(ctx *fiber.Ctx, methodName string, methodHandler methodJsonRPC) (err error) {

	span := otg.SpanFromContext(ctx.UserContext())
	span.SetTag("method", methodName)

	methodHTTP := ctx.Method()
	if methodHTTP != fiber.MethodPost {
		ext.Error.Set(span, true)
		span.SetTag("msg", "only POST method supported")
		ctx.Response().SetStatusCode(fiber.StatusMethodNotAllowed)
		if _, err = ctx.WriteString("only POST method supported"); err != nil {
			return
		}
	}
	var request baseJsonRPC
	var response *baseJsonRPC
	if err = json.Unmarshal(ctx.Body(), &request); err != nil {
		ext.Error.Set(span, true)
		span.SetTag("msg", "request body could not be decoded: "+err.Error())
		return sendResponse(ctx, makeErrorResponseJsonRPC([]byte("\"0\""), parseError, "request body could not be decoded: "+err.Error(), nil))
	}
	methodNameOrigin := request.Method
	method := strings.ToLower(request.Method)
	if method != "" && method != methodName {
		ext.Error.Set(span, true)
		span.SetTag("msg", "invalid method "+methodNameOrigin)
		return sendResponse(ctx, makeErrorResponseJsonRPC(request.ID, methodNotFoundError, "invalid method "+methodNameOrigin, nil))
	}
	response = methodHandler(ctx.UserContext(), request)
	if response != nil {
		return sendResponse(ctx, response)
	}
	return
}
