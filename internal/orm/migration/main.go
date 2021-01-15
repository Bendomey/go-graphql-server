package migration

import (
	"fmt"

	log "github.com/Bendomey/graphql-boilerplate/internal/logger"

	"github.com/Bendomey/graphql-boilerplate/internal/orm/migration/jobs"
	"github.com/Bendomey/graphql-boilerplate/internal/orm/models"
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func updateMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
	)
	return err
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Info("[Migration.InitSchema] Initializing database schema")
		db.Exec("create extension \"uuid-ossp\";")
		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}
		// Add more jobs, etc here
		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers,
	})
	return m.Migrate()
}

// # #!/bin/bash
// # printf "\nRegenerating gqlgen files\n"
// # rm -f internal/gql/generated.go \
// #     internal/gql/models/generated.go \
// #     internal/gql/resolvers/generated/resolver.go
// # time go run -v github.com/99designs/gqlgen generate
// # printf "\nDone.\n\n"
