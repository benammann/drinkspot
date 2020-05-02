package router

import (
	"github.com/benammann/drinkspot-core/api/utility"
	"github.com/gin-gonic/gin"
)

func middleware(router *Router) {

	router.engine.Use(
		middlewareSetMemoryEnvironment(router),
	)

}

func middlewareSetMemoryEnvironment(router *Router) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		utility.SetEnvironment(ctx, router.env)
	}
}
