package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/99designs/gqlgen/example/starwars/generated"
	"github.com/99designs/gqlgen/graphql"

	"github.com/Bendomey/graphql-boilerplate/internal/gql"
)

type Resolver struct{}

// NewGraphqlServer creates a graphql server of all microservices
func NewGraphqlServer() (*Resolver, error) {
	// connect to user service
	// userService := services.NewUserService(repo)
	return &Resolver{}, nil
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// ToExecutableSchema retrieves services here
func (r *Resolver) ToExecutableSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: nil,
	})
}
