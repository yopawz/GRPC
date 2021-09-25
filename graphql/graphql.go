package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "World", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)
	fmt.Println(err)
	// HandleError("Cannot Make New Schema", err)

	query := `{
		hello
	}`

	params := graphql.Params{Schema: schema, RequestString: query}

	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operations, error: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)

	fmt.Printf("%s \n", rJSON)
}
