// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/seniorGolang/tg/v2/example/implement"
)

func (http *httpUser) getUser(ctx context.Context, request requestUserGetUser) (response responseUserGetUser, err error) {
	span := opentracing.SpanFromContext(ctx)
	response.User, err = http.svc.GetUser(ctx, request.Cookie, request.UserAgent)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		errData := toString(err)
		ext.Error.Set(span, true)
		span.SetTag("msg", err.Error())
		if errData != "{}" {
			span.SetTag("errData", errData)
		}
	}
	return
}
func (http *httpUser) serveGetUser(ctx *fiber.Ctx) (err error) {
	span := extractSpan(http.log, fmt.Sprintf("request:%s", ctx.Path()), ctx)
	defer injectSpan(http.log, span, ctx)
	defer span.Finish()
	if value := ctx.Context().Value(CtxCancelRequest); value != nil {
		ext.Error.Set(span, true)
		span.SetTag("msg", "request canceled")
		return
	}
	var request requestUserGetUser
	ctx.Response().SetStatusCode(204)

	if _userAgent := string(ctx.Request().Header.Peek("User-Agent")); _userAgent != "" {
		var userAgent string
		userAgent = _userAgent
		request.UserAgent = userAgent
	}

	if _cookie := ctx.Cookies("sessionCookie"); _cookie != "" {
		var cookie string
		cookie = _cookie
		request.Cookie = cookie
	}

	var response responseUserGetUser
	methodContext := opentracing.ContextWithSpan(ctx.Context(), span)
	if response, err = http.getUser(methodContext, request); err == nil {
		return sendResponse(http.log, ctx, response)
	}
	if errCoder, ok := err.(withErrorCode); ok {
		ctx.Status(errCoder.Code())
	} else {
		ctx.Status(fiber.StatusInternalServerError)
	}
	return sendResponse(http.log, ctx, err)
}
func (http *httpUser) uploadFile(ctx context.Context, request requestUserUploadFile) (response responseUserUploadFile, err error) {
	span := opentracing.SpanFromContext(ctx)
	err = http.svc.UploadFile(ctx, request.FileBytes)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		errData := toString(err)
		ext.Error.Set(span, true)
		span.SetTag("msg", err.Error())
		if errData != "{}" {
			span.SetTag("errData", errData)
		}
	}
	return
}
func (http *httpUser) serveUploadFile(ctx *fiber.Ctx) (err error) {
	span := extractSpan(http.log, fmt.Sprintf("request:%s", ctx.Path()), ctx)
	defer injectSpan(http.log, span, ctx)
	defer span.Finish()
	if value := ctx.Context().Value(CtxCancelRequest); value != nil {
		ext.Error.Set(span, true)
		span.SetTag("msg", "request canceled")
		return
	}
	var request requestUserUploadFile

	if request.FileBytes, err = uploadFile(ctx, "fileBytes"); err != nil {
		ext.Error.Set(span, true)
		span.SetTag("msg", "upload file 'fileBytes' error: "+err.Error())
		ctx.Status(fiber.StatusBadRequest)
		return sendResponse(http.log, ctx, "upload file 'fileBytes' error: "+err.Error())
	}
	var response responseUserUploadFile
	methodContext := opentracing.ContextWithSpan(ctx.Context(), span)
	if response, err = http.uploadFile(methodContext, request); err == nil {
		return sendResponse(http.log, ctx, response)
	}
	if errCoder, ok := err.(withErrorCode); ok {
		ctx.Status(errCoder.Code())
	} else {
		ctx.Status(fiber.StatusInternalServerError)
	}
	return sendResponse(http.log, ctx, err)
}
func (http *httpUser) customResponse(ctx context.Context, request requestUserCustomResponse) (response responseUserCustomResponse, err error) {
	span := opentracing.SpanFromContext(ctx)
	err = http.svc.CustomResponse(ctx, request.Arg0, request.Arg1, request.Opts...)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		errData := toString(err)
		ext.Error.Set(span, true)
		span.SetTag("msg", err.Error())
		if errData != "{}" {
			span.SetTag("errData", errData)
		}
	}
	return
}
func (http *httpUser) serveCustomResponse(ctx *fiber.Ctx) (err error) {
	span := extractSpan(http.log, fmt.Sprintf("request:%s", ctx.Path()), ctx)
	defer injectSpan(http.log, span, ctx)
	defer span.Finish()
	if value := ctx.Context().Value(CtxCancelRequest); value != nil {
		ext.Error.Set(span, true)
		span.SetTag("msg", "request canceled")
		return
	}
	var request requestUserCustomResponse
	if err = json.Unmarshal(ctx.Request().Body(), &request); err != nil {
		ext.Error.Set(span, true)
		span.SetTag("msg", "request body could not be decoded: "+err.Error())
		ctx.Response().SetStatusCode(fiber.StatusBadRequest)
		_, err = ctx.WriteString("request body could not be decoded: " + err.Error())
		return
	}
	return implement.CustomResponseHandler(ctx, http.base, request.Arg0, request.Arg1, request.Opts...)
}
func (http *httpUser) customHandler(ctx context.Context, request requestUserCustomHandler) (response responseUserCustomHandler, err error) {
	span := opentracing.SpanFromContext(ctx)
	err = http.svc.CustomHandler(ctx, request.Arg0, request.Arg1, request.Opts...)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		errData := toString(err)
		ext.Error.Set(span, true)
		span.SetTag("msg", err.Error())
		if errData != "{}" {
			span.SetTag("errData", errData)
		}
	}
	return
}
func (http *httpUser) serveCustomHandler(ctx *fiber.Ctx) (err error) {
	span := extractSpan(http.log, fmt.Sprintf("request:%s", ctx.Path()), ctx)
	defer injectSpan(http.log, span, ctx)
	defer span.Finish()
	if value := ctx.Context().Value(CtxCancelRequest); value != nil {
		ext.Error.Set(span, true)
		span.SetTag("msg", "request canceled")
		return
	}
	var request requestUserCustomHandler
	if err = json.Unmarshal(ctx.Request().Body(), &request); err != nil {
		ext.Error.Set(span, true)
		span.SetTag("msg", "request body could not be decoded: "+err.Error())
		ctx.Response().SetStatusCode(fiber.StatusBadRequest)
		_, err = ctx.WriteString("request body could not be decoded: " + err.Error())
		return
	}

	var response responseUserCustomHandler
	methodContext := opentracing.ContextWithSpan(ctx.Context(), span)
	if response, err = http.customHandler(methodContext, request); err == nil {
		return sendResponse(http.log, ctx, response)
	}
	if errCoder, ok := err.(withErrorCode); ok {
		ctx.Status(errCoder.Code())
	} else {
		ctx.Status(fiber.StatusInternalServerError)
	}
	return sendResponse(http.log, ctx, err)
}
