package wikigraphql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// ExecuteQuery executes a query with the given schema
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	fmt.Println("ExecuteQuery: ", query)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors inside ExecuteQuery: %v", result.Errors)
	}

	return result
}
