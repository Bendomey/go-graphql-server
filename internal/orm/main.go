// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package orm

import (
	"fmt"

	log "github.com/Bendomey/graphql-boilerplate/internal/logger"

	"github.com/Bendomey/graphql-boilerplate/internal/orm/migration"

	"github.com/Bendomey/graphql-boilerplate/pkg/utils"
	//Imports the database dialect of choice
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var autoMigrate, seedDB bool
var host, user, password, dbname, port, sslmode string

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

func init() {
	host = utils.MustGet("GORM_HOST")
	user = utils.MustGet("GORM_USER")
	password = utils.MustGet("GORM_PASSWORD")
	dbname = utils.MustGet("GORM_DBNAME")
	port = utils.MustGet("GORM_PORT")
	seedDB = utils.MustGetBool("GORM_SEED_DB")
	autoMigrate = utils.MustGetBool("GORM_AUTOMIGRATE")
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)), &gorm.Config{})
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}

	// Automigrate tables
	if autoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	log.Info("[ORM] Database connection initialized.")
	return orm, err
}
