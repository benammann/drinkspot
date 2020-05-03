package router

import (
	"fmt"
	"github.com/benammann/drinkspot-core/api/app/controller"
	"github.com/benammann/drinkspot-core/api/memory"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Router struct {
	env    *memory.Environment
	engine *gin.Engine
	port   string
}

// Creates a new router instance
func NewRouter(env *memory.Environment) *Router {

	// creates a new router instance and sets it's defaults
	instance := &Router{

		// holds the global enviroment like database connection etc
		env: env,

		// holds the gin instance
		engine: gin.Default(),

		// holds the port. Format :$PORT
		port: fmt.Sprintf(":%s", os.Getenv("API_PORT")),
	}

	instance.engine.Use(cors.Default())

	// sets up the middleware
	middleware(instance)

	// adds the /graphql endpoint
	graphQLEndpoint(instance)

	// adds the playground endpoint
	playgroundEndpoint(instance)

	// add the authentication endpoints
	controller.Authentication(instance.engine)

	// return the router instance
	return instance
}

// starts the http router
func (router *Router) ListenAndServe() error {
	if gin.IsDebugging() {
		return router.listenAndServeHttp()
	} else {
		return router.listenAndServeHttps()
	}
}

// starts the router server using HTTP
func (router *Router) listenAndServeHttp() error {
	return http.ListenAndServe(router.port, router.engine)
}

// stars the router using HTTPS
// make sure to set API_CERT_FILE and API_KEY_FILE inside your .env file
// is implemented for production use
func (router *Router) listenAndServeHttps() error {
	return http.ListenAndServeTLS(router.port, os.Getenv("API_CERT_FILE"), os.Getenv("API_KEY_FILE"), router.engine)
}
