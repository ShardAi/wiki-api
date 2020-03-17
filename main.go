package main

import (
	"fmt"
	"log"
	"net/http"

	"wiki-api/sql"
	"wiki-api/wikigraphql"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	schema, db := initializeAPI()
	defer db.Close()

	h := handler.New(&handler.Config{
		Schema: schema,
		Pretty: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func initializeAPI() (*graphql.Schema, *sql.Db) {

	db, err := sql.New(
		sql.ConnString("localhost\\localDBInstance", 53920),
	)
	if err != nil {
		log.Fatal(err)
	}

	queryRoot, mutationRoot := wikigraphql.NewRoot(db)

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: queryRoot.Query, Mutation: mutationRoot.Mutation},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	return &schema, db
}
