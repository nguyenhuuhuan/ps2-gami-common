package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"net/http"
)

// Tracing - distribute tracing with OpenCensus using Google StackTraceDriver
func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		var span *trace.Span

		path := c.Request.URL.Path

		ctx, span = trace.StartSpan(ctx, path)
		span.AddAttributes(requestAttrs(c.Request)...)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func requestAttrs(r *http.Request) []trace.Attribute {
	return []trace.Attribute{
		trace.StringAttribute(ochttp.PathAttribute, r.URL.Path),
		trace.StringAttribute(ochttp.HostAttribute, r.URL.Host),
		trace.StringAttribute(ochttp.MethodAttribute, r.Method),
		trace.StringAttribute(ochttp.UserAgentAttribute, r.UserAgent()),
	}
}
