package tcp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
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

func logWithReqID(reqID string, format string, args ...any) {
	message := fmt.Sprintf("[%s] ", reqID)
	message += fmt.Sprintf(format, args...)

	log.Print(message)
}

func (s *Server) handle(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(s.cfg.Timeout))
	conn.SetWriteDeadline(time.Now().Add(s.cfg.Timeout))
	reqID := uuid.NewString()

	defer closeConnection(reqID, conn)

	logWithReqID(reqID, "start to handle connection")

	ctx := context.Background()

	challenge, err := s.challenger.Generate()
	if err != nil {
		logWithReqID(reqID, "failed to generate challenge: `%s`", err)

		return
	}

	cr := challengeResponse{Challenge: challenge}
	if err = json.NewEncoder(conn).Encode(&cr); err != nil {
		logWithReqID(reqID, "send challenge: `%s`", err)

		return
	}

	logWithReqID(reqID, "challenge sent: `%s`", cr)

	sr := solutionRequest{}
	if err = json.NewDecoder(conn).Decode(&sr); err != nil {
		logWithReqID(reqID, "failed get solution: %s", err)

		return
	}

	err = s.challenger.Verify(challenge, sr.Solution)
	if err != nil {
		logWithReqID(reqID, "failed to verify PoW solution: %s", err)

		return
	}

	quote, err := s.storage.Quote(ctx)
	if err != nil {
		logWithReqID(reqID, "failed to get quote: `%s`", err)

		return
	}

	qr := quoteResponse{Quote: quote}
	if err = json.NewEncoder(conn).Encode(qr); err != nil {
		logWithReqID(reqID, "failed to send quote response: `%s`", err)

		return
	}

	logWithReqID(reqID, "quote sent: `%s`", qr)
}

func closeConnection(reqID string, conn net.Conn) {
	if err := conn.Close(); err != nil {
		logWithReqID(reqID, "failed to close connection: `%s`", err)

		return
	}

	logWithReqID(reqID, "connection is closed")
}
