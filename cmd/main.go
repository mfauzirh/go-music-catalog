package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-music-catalog/internal/configs"
	membershipsHandler "github.com/mfauzirh/go-music-catalog/internal/handler/memberships"
	"github.com/mfauzirh/go-music-catalog/internal/models/memberships"
	membershipsRepository "github.com/mfauzirh/go-music-catalog/internal/repository/memberships"
	membershipsService "github.com/mfauzirh/go-music-catalog/internal/service/memberships"
	"github.com/mfauzirh/go-music-catalog/pkg/internalsql"
	"github.com/rs/zerolog/log"
)

func main() {
	var cfg *configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initalize config")
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	membershipRepo := membershipsRepository.NewRepository(db)
	membershipSvc := membershipsService.NewService(cfg, membershipRepo)
	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
