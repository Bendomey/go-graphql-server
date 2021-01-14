package server

import (
	"log"

	"github.com/Bendomey/graphql-boilerplate/internal/handlers"
	"github.com/gin-gonic/gin"
)

var HOST, PORT string

func init() {
	HOST = "localhost"
	PORT = "7777"
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + HOST + ":" + PORT)
	log.Fatalln(r.Run(HOST + ":" + PORT))
}
