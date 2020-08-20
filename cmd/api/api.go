package main

import (
	"coding-challenge-go/pkg/api"
	"os"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	engine, err := api.CreateAPIEngine()

	if err != nil {
		log.Error().Err(err).Msg("Fail to create server")
	}

	log.Info().Msg("Start server")
	log.Fatal().Err(engine.Run(os.Getenv("LISTEN"))).Msg("Fail to listen and serve")
}
