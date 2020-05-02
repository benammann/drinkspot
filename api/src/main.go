package main

import (
	"fmt"
	"github.com/benammann/drinkspot-core/api/database"
	"github.com/benammann/drinkspot-core/api/memory"
	"github.com/benammann/drinkspot-core/api/router"
)

func main() {

	fmt.Println("hi")

	db := createDatabaseConnection()
	defer db.Connection.Close()

	environment := &memory.Environment{
		Db: db,
	}

	// starts the api router
	startRouter(environment)

}

// creates the database connection
func createDatabaseConnection() *database.Database {
	return database.NewDatabase()
}

// initializes and starts the http router
func startRouter(env *memory.Environment) {

	routerInstance := router.NewRouter(env)
	err := routerInstance.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("Could not start HTTP Server: %s", err.Error()))
	}

}
