package main

import (
	"log"

	"github.com/Tsapen/wow/internal/client"
	"github.com/Tsapen/wow/internal/config"
	"github.com/Tsapen/wow/internal/solver"
)

func main() {
	cfg, err := config.GetForClient()
	if err != nil {
		log.Fatal("failed to read config: ", err)
	}

	solver := solver.New()

	client := client.New(cfg, solver)

	quote, err := client.GetQuote()
	if err != nil {
		log.Fatal("failed to get quote from server: ", err)
	}

	log.Printf("success! quote is \"%s\"\n", quote)
}
