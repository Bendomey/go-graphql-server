package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Bendomey/graphql-boilerplate/internal/gql"
	"github.com/Bendomey/graphql-boilerplate/internal/gql/models"
)

type Resolver struct{}

// NewGraphqlServer creates a graphql server of all microservices
func NewGraphqlServer() (*Resolver, error) {
	// connect to user service
	// userService := services.NewUserService(repo)

	return &Resolver{}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context, id *string) (*models.Users, error) {
	panic("not implemented")
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// ToExecutableSchema retrieves services here
func (r *Resolver) ToExecutableSchema() graphql.ExecutableSchema {
	return gql.NewExecutableSchema(gql.Config{
		Resolvers: nil,
	})
}
