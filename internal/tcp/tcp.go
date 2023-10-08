package tcp

import (
	"fmt"
	"log"
	"net"

	"github.com/Tsapen/wow/internal/config"
	"github.com/Tsapen/wow/internal/wow"
)

type Server struct {
	cfg        *config.ServerConfig
	listener   net.Listener
	challenger wow.Challenger
	storage    wow.Storage

	sem chan struct{}
}

// NewServer creates new tcp server.
func NewServer(cfg *config.ServerConfig, challenger wow.Challenger, storage wow.Storage) (*Server, error) {
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return nil, fmt.Errorf("make tcp listener: %w", err)
	}

	sem := make(chan struct{}, cfg.ConnectionMaxCount)
	for i := 0; i < cfg.ConnectionMaxCount; i++ {
		sem <- struct{}{}
	}

	return &Server{
		cfg:        cfg,
		listener:   listener,
		challenger: challenger,
		storage:    storage,

		sem: sem,
	}, nil
}

// ListenAndServe listens and serves tcp connections.
func (s *Server) ListenAndServe() {
	defer s.listener.Close()

	log.Printf("Server started on %s\n", s.cfg.Address)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print("failed to accept connection:", err.Error())

			return
		}

		go func() {
			<-s.sem

			s.handle(conn)

			s.sem <- struct{}{}
		}()
	}
}
