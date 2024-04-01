package tracing

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type httpTrace struct {
}

func Http() *httpTrace {
	return &httpTrace{}
}

type HttpExtractOption http.Header

func (h HttpExtractOption) parse() interface{} {
	return nil
}

func (httpTrace) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从消息头获取父级的上下文
		span := StartSpan(ctx, ctx.FullPath(), HttpExtractOption(ctx.Request.Header))
		defer span.End()
	}
}
