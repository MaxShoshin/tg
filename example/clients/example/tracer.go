// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package example

import (
	"context"
	"github.com/gofiber/fiber/v2"
	otg "github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func extractSpan(ctx context.Context, opName string) (span otg.Span) {

	var opts []otg.StartSpanOption
	span = otg.SpanFromContext(ctx)

	if span == nil {
		log.Ctx(ctx).Debug().Msg("context does not contain span")
	} else {
		opts = append(opts, otg.ChildOf(span.Context()))
	}

	span = otg.GlobalTracer().StartSpan(opName, opts...)
	return
}

func injectSpan(span otg.Span, request *fiber.Request) {
	headers := make(http.Header)
	if err := otg.GlobalTracer().Inject(span.Context(), otg.HTTPHeaders, otg.HTTPHeadersCarrier(headers)); err != nil {
		log.Warn().Err(err).Msg("inject span to HTTP headers")
	}
	for key, values := range headers {
		request.Header.Set(key, strings.Join(values, ";"))
	}
}
