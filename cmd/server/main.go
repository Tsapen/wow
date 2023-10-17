package main

import (
	"github.com/rs/zerolog/log"

	"github.com/Tsapen/wow/internal/challenger"
	"github.com/Tsapen/wow/internal/config"
	"github.com/Tsapen/wow/internal/solver"
	"github.com/Tsapen/wow/internal/storage"
	"github.com/Tsapen/wow/internal/tcp"
)

type Config struct {
	Address string
}

func main() {
	cfg, err := config.GetForServer()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	solverService := solver.New()

	challengerService := challenger.New(solverService)

	st := storage.New()

	server, err := tcp.NewServer(cfg, challengerService, st)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run tcp server")
	}

	server.ListenAndServe()
}
