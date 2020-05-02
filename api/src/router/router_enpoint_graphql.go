package router

import (
	middleware2 "github.com/benammann/drinkspot-core/api/app/middleware"
	"github.com/benammann/drinkspot-core/api/graphql"
)

func graphQLEndpoint(router *Router) {

	graphQlGroup := router.engine.Group("/graphql")
	graphQlGroup.Use(middleware2.AuthenticateUser)

	graphQlGroup.POST("/", graphql.HttpHandler)

}
