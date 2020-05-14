package gin_graphql

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var GinCtxKey struct{}
// var GinCtxKey "ginkey"

var GinCtxKey string = "ginkey"

func FromStd(handler http.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), GinCtxKey, ctx))
		handler(ctx.Writer, request)
	}
}
func GetGinCtxFromStdCtx(ctx context.Context) *gin.Context {
	return ctx.Value(GinCtxKey).(*gin.Context)
}
