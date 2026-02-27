package database

import (
	"context"
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"

	"github.com/chokoskoder/dashboard-script/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.json 
var migrations embed.FS
//this is to cd into the different directories ??
//what is this ??

type MigrationAction string

const (
	ActionUp   MigrationAction = "up"
	ActionDown MigrationAction = "down" // This usually means "Down 1 step"
)

func Migrate(ctx context.Context , logger *slog.Logger , cfg *config.Config , action MigrationAction) error {
	//start by connecting to the db

	//start a new migrator using go.migrator
	//get the migration subtree so far
	//lod the migration subtree
	//get current version
	//migrate

	//setup the soure driver , to look into the migrations directory 
		sourceDriver , err := iofs.New(migrations , "migrations")
	if err != nil{
		return fmt.Errorf("failed to create iofs source driver : %w" , err)
	}

	//create the migrator instance
	m , err := migrate.NewWithSourceInstance(
		"iofs",
		sourceDriver,
		cfg.DBURI,
	)

	var migrateErr error
	switch action {
	case ActionUp:
		migrateErr = m.Up()
	case ActionDown:
		migrateErr = m.Steps(-1)
	default:
		return fmt.Errorf("unknown migration action: %s", action)
	}

	if migrateErr != nil {
		if errors.Is(migrateErr, migrate.ErrNoChange) {
			//add logging here
			return nil
		}
		return fmt.Errorf("migration failed: %w", migrateErr)
	}
	return nil
}