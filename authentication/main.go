package main

import (
	"auth/api"
	db "auth/db/sqlc"
	"auth/pkgs/logger"
	"auth/pkgs/tools"
	"auth/utils"
	"context"
	"database/sql"
	"os"
	"time"
)

func main() {

	// Create the logger.
	logger := logger.New(os.Stdout, logger.LevelInfo)

	// load the env vars
	config, err := utils.LoadENVFromFile(".", "app", "env")
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	// connect to db
	pdb, err := tools.Connect("postgres", 10, 1*time.Second, func() (*sql.DB, error) {
		return db.Connect(db.ContectConfig{
			DataSourceName:  config.DATABASE_SOURCE_NAME,
			MaxOpenConns:    config.DB_MAX_OPEN_CONNECTION,
			MaxIdleConns:    config.DB_MAX_IDLE_CONNECTION,
			ConnMaxIdleTime: config.ACCESS_TOKEN_DURATION,
			DriverName:      "postgres",
		})
	})
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintInfo("Connected the database", nil)
	defer pdb.Close()

	// run migration aganist the db
	err = db.RunMigrations(config.DATABASE_SOURCE_NAME, config.MIGRATION_URL)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	logger.PrintInfo("Migrarated successfully", nil)

	store := db.New(pdb)

	store.GetUserByID(context.Background(), 1)
	// initialize the server and start it
	srv, err := api.NewServer(&api.NewServerArgs{
		Logger: logger,
		Config: config,
	})
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintFatal(srv.Start(config.PORT), nil)
}
