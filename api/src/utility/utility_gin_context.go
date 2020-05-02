package utility

import (
	"context"
	"github.com/gin-gonic/gin"
)

func GetGinContext(ctx context.Context) *gin.Context {
	return ctx.Value("gin_context").(*gin.Context)
}

func GinRichMutation(localCtx context.Context, localArgs interface{}, resolver func(ctx *gin.Context, args interface{}) interface{}) interface{} {
	ginContext := GetGinContext(localCtx)
	return resolver(ginContext, localArgs)
}

func GinRichQuery(localCtx context.Context, resolver func(ctx *gin.Context) (interface{}, error)) (interface{}, error) {
	ginContext := GetGinContext(localCtx)
	return resolver(ginContext)
}
