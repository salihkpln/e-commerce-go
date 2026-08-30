package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salihkpln/e-commerce-go/internal/config"
	"github.com/salihkpln/e-commerce-go/internal/database"
	"github.com/salihkpln/e-commerce-go/internal/logger"
)

func main() {
	log := logger.New()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	db, err := database.NewDatabaseConnection(&cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	mainDb, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	defer func() {
		if err := mainDb.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close the database connection")
		}
	}()
	gin.SetMode(cfg.Server.GinMode)

	log.Info().Msg("Starting server")
}
