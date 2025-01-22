package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-music-catalog/internal/configs"
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

	_, err = internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	r := gin.Default()

	r.Run(cfg.Service.Port)
}
