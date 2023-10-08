package main

import (
	"log"

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
		log.Fatal("failed to read config: ", err)
	}

	solver := solver.New()

	challenger := challenger.New(solver)

	storage := storage.New()

	server, err := tcp.NewServer(cfg, challenger, storage)
	if err != nil {
		log.Fatal("failed to run tcp server: ", err)
	}

	server.ListenAndServe()
}
