package wikigraphql

import "github.com/graphql-go/graphql"

// WikiPage parser type object for the wikipage to be used in parsing graphql queries
var WikiPage = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WikiPage",
		Fields: graphql.Fields{
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Tags": &graphql.Field{
				Type: graphql.String,
			},
			"Ingress": &graphql.Field{
				Type: graphql.String,
			},
			"MainText": &graphql.Field{
				Type: graphql.String,
			},
			"SideBarInfo": &graphql.Field{
				Type: graphql.String,
			},
			"ProfileImagePath": &graphql.Field{
				Type: graphql.String,
			},
			"BodyImagePath": &graphql.Field{
				Type: graphql.String,
			},
			"Visible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Author": &graphql.Field{
				Type: graphql.String,
			},
			"LastChanged": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
