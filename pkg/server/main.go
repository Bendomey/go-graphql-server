package server

import (
	"log"

	"github.com/Bendomey/graphql-boilerplate/internal/handlers"
	"github.com/Bendomey/graphql-boilerplate/pkg/utils"
	"github.com/gin-gonic/gin"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	port = utils.MustGet("PORT")
	gqlPath = utils.MustGet("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED")
}

// Run web server
func Run() {
	r := gin.Default()
	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPath, handlers.PlaygroundHandler(gqlPgPath))
		log.Println("GraphQL Playground @ " + gqlPgPath)
	}
	r.POST(gqlPath, handlers.GraphqlHandler())
	// log.Println("Running @ http://" + host + ":" + port)
	log.Println("Running on port: " + port)
	log.Fatalln(r.Run())
}
