package tcp

import (
	"context"
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

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

	log.Info().Msgf("Server started on %s", s.cfg.Address)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Warn().Err(err).Msgf("failed to accept connection")

			return
		}

		go func() {
			<-s.sem

			ctx := context.Background()
			reqID := uuid.NewString()
			ctx = wow.WithReqID(ctx, reqID)
			if err := s.handle(ctx, conn); err != nil {
				log.Info().Str("request_id", reqID).Err(err).Msgf("failed to handle connection")

				return
			}

			s.sem <- struct{}{}
		}()
	}
}
