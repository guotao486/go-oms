/*
 * @Author: GG
 * @Date: 2023-02-19 11:15:59
 * @LastEditTime: 2023-02-19 11:54:47
 * @LastEditors: GG
 * @Description: tracer 中间件
 * @FilePath: \oms\internal\middleware\tracer.go
 *
 */
package middleware

import (
	"context"
	"oms/global"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)

		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(), // 上下文
				global.Tracer,       // jager tracer 对象
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}

		defer span.Finish()

		// 获取traceID 和 spanID
		var traceID string
		var spanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext := spanContext.(jaeger.SpanContext)
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
		}

		c.Set("X-Trace-ID", traceID)
		c.Set("X-Trace-ID", spanID)
		// 获取traceID 和 spanID end
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
