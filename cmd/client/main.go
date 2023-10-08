package main

import (
	"github.com/rs/zerolog/log"

	"github.com/Tsapen/wow/internal/client"
	"github.com/Tsapen/wow/internal/config"
	"github.com/Tsapen/wow/internal/solver"
)

func main() {
	cfg, err := config.GetForClient()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	solver := solver.New()

	client := client.New(cfg, solver)

	quote, err := client.GetQuote()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get quote from server")
	}

	log.Info().Msgf("success! quote is \"%s\"\n", quote)
}
