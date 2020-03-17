package wikigraphql

import (
	"wiki-api/sql"

	"github.com/graphql-go/graphql"
)

// QRoot structure keeping the graphql object to handle queries
type QRoot struct {
	Query *graphql.Object
}

// MRoot structure keeping graphql object to handle mutations
type MRoot struct {
	Mutation *graphql.Object
}

// NewRoot Creates and returns a new Root structure
func NewRoot(db *sql.Db) (*QRoot, *MRoot) {
	queryResolver := QueryResolver{db: db}
	mutationResolver := MutationResolver{db: db}

	queryRoot := QRoot{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootQuery",
				Fields: graphql.Fields{
					"GetPages": &graphql.Field{
						Type: graphql.NewList(WikiPage),
						Args: graphql.FieldConfigArgument{
							"Title": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"Tags": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: queryResolver.GetPagesResolver,
					},
				},
			},
		),
	}
	mutRoot := MRoot{
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootMutation",
				Fields: graphql.Fields{
					"SavePage": &graphql.Field{
						Type: WikiPage,
						Args: graphql.FieldConfigArgument{
							"Title": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"Tags": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"Ingress": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"MainText": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"SideBarInfo": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"ProfileImagePath": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"BodyImagePath": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"Visible": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Boolean),
							},
							"Author": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: mutationResolver.SavePageResolver,
					},
				},
			},
		),
	}
	return &queryRoot, &mutRoot
}
