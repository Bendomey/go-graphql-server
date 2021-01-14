package server

import (
	"log"

	"github.com/Bendomey/graphql-boilerplate/internal/handlers"
	"github.com/Bendomey/graphql-boilerplate/pkg/utils"
	"github.com/gin-gonic/gin"
)

var port string

func init() {
	port = utils.MustGet("PORT")
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	// log.Println("Running @ http://" + host + ":" + port)
	log.Println("Running on port: " + port)
	log.Fatalln(r.Run())
}
