package graphql

import (
	api_current_user "github.com/benammann/drinkspot-core/api/graphql/resources/api-current-user"
	api_drink_spot "github.com/benammann/drinkspot-core/api/graphql/resources/api-drink-spot"
	"github.com/benammann/drinkspot-core/api/graphql/resources/api-version"
	"github.com/graph-gophers/graphql-go"
	"strings"
)

type Resolver struct{}

var Schema = `

	type Query {}
	type Mutation {}

	schema {
		query: Query
		mutation: Mutation
	}

`

func BuildSchemaDefinition() string {

	resources := []string{

		// schema root first
		Schema,

		// then all the modules
		api_current_user.Schema,
		api_version.Schema,
		api_drink_spot.Schema,
	}

	return strings.Join(resources[:], "")

}

func getSchema() *graphql.Schema {

	// build the schema definition
	definition := BuildSchemaDefinition()

	// then return the graphql struct
	return graphql.MustParseSchema(definition, &Resolver{})
}
