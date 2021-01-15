package handlers

import (
	"github.com/99designs/gqlgen/handler"
	resolvers "github.com/Bendomey/graphql-boilerplate/internal/gql/resolvers"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(srv *resolvers.Resolver) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// c := gql.Config{
	// 	Resolvers: &resolvers.Resolver{},
	// }

	h := handler.GraphQL(srv.ToExecutableSchema())

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Boilerplate Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
