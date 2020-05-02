package router

import (
	"github.com/friendsofgo/graphiql"
	"github.com/gin-gonic/gin"
)

func playgroundEndpoint(router *Router) {
	if gin.IsDebugging() {

		// add the dashboard
		dashboard, err := graphiql.NewGraphiqlHandler("/graphql")
		if err != nil {
			panic(err)
		}

		router.engine.GET("/playground", func(context *gin.Context) {
			dashboard.ServeHTTP(context.Writer, context.Request)
		})

	}
}
