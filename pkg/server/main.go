package server

import (
	"log"

	"github.com/Bendomey/graphql-boilerplate/internal/handlers"
	"github.com/Bendomey/graphql-boilerplate/pkg/utils"
	"github.com/gin-gonic/gin"
)

var host, port string

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("PORT")
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(r.Run(host + ":" + port))
}
