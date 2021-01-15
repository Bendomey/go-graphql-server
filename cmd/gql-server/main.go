package main

import (
	log "github.com/Bendomey/graphql-boilerplate/internal/logger"
	"github.com/Bendomey/graphql-boilerplate/internal/orm"
	"github.com/Bendomey/graphql-boilerplate/pkg/server"
)

func main() {
	// Create a new ORM instance to send it to our
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm)
}
