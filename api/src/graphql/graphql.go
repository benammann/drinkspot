package graphql

import (
	api_current_user "github.com/benammann/drinkspot-core/api/graphql/resources/api-current-user"
	"github.com/benammann/drinkspot-core/api/graphql/resources/api-version"
	"github.com/graph-gophers/graphql-go"
	"strings"
)

type Resolver struct{}

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	type Query {}
	type Mutation {}
`

func BuildSchemaDefinition() string {

	resources := []string{

		// schema root first
		Schema,

		// then all the modules
		api_current_user.Schema,
		api_version.Schema,
	}

	return strings.Join(resources[:], "")

}

func getSchema() *graphql.Schema {

	// build the schema definition
	definition := BuildSchemaDefinition()

	// then return the graphql struct
	return graphql.MustParseSchema(definition, &Resolver{})
}
