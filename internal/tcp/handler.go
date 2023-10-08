package tcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/Tsapen/wow/internal/wow"
)

type challengeResponse struct {
	Challenge string `json:"challenge"`
}

type solutionRequest struct {
	Solution string `json:"solution"`
}

type quoteResponse struct {
	Quote string `json:"quote"`
}

func (s *Server) handle(ctx context.Context, conn net.Conn) (err error) {
	conn.SetReadDeadline(time.Now().Add(s.cfg.Timeout))
	conn.SetWriteDeadline(time.Now().Add(s.cfg.Timeout))

	logger := log.With().Str("request_id", wow.ReqIDFromCtx(ctx)).Logger()

	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			err = wow.HandleErrPair(fmt.Errorf("close error: %w", closeErr), err)

		} else {
			logger.Info().Msg("connection closed")
		}
	}()

	logger.Info().Msg("start to handle connection")

	challenge, err := s.challenger.Generate()
	if err != nil {
		return fmt.Errorf("failed to generate challenge: %w", err)
	}

	cr := challengeResponse{Challenge: challenge}
	if err = json.NewEncoder(conn).Encode(&cr); err != nil {
		return fmt.Errorf("send challenge: %w", err)
	}

	logger.Info().Any("request", cr).Msg("challenge sent")

	sr := solutionRequest{}
	if err = json.NewDecoder(conn).Decode(&sr); err != nil {
		return fmt.Errorf("failed get solution: %w", err)
	}

	err = s.challenger.Verify(challenge, sr.Solution)
	if err != nil {
		return fmt.Errorf("failed to verify PoW solution: %w", err)
	}

	quote, err := s.storage.Quote(ctx)
	if err != nil {
		return fmt.Errorf("failed to get quote: %w", err)
	}

	qr := quoteResponse{Quote: quote}
	if err = json.NewEncoder(conn).Encode(qr); err != nil {
		return fmt.Errorf("failed to send quote response: %w`", err)
	}

	logger.Info().Any("request", qr).Msg("quote sent")

	return nil
}
