package client

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/Tsapen/wow/internal/config"
	"github.com/Tsapen/wow/internal/wow"
)

type Client struct {
	cfg    *config.ClientConfig
	solver wow.Solver
}

type challengeResponse struct {
	Challenge string `json:"challenge"`
}

type solutionRequest struct {
	Solution string `json:"solution"`
}

type quoteResponse struct {
	Quote string `json:"quote"`
}

func New(cfg *config.ClientConfig, solver wow.Solver) *Client {
	return &Client{
		cfg:    cfg,
		solver: solver,
	}
}

func (c *Client) GetQuote() (quote string, err error) {
	conn, err := net.Dial("tcp", c.cfg.Address)
	if err != nil {
		return "", fmt.Errorf("connect to the server: %w", err)
	}

	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			err = wow.HandleErrPair(fmt.Errorf("close connection: %w", closeErr), err)
		}
	}()

	cr := challengeResponse{}
	if err = json.NewDecoder(conn).Decode(&cr); err != nil {
		return "", fmt.Errorf("get challenge: %w", err)
	}

	solution, err := c.solver.Solve(cr.Challenge)
	if err != nil {
		return "", fmt.Errorf("solve PoW challenge: %w", err)
	}

	sr := solutionRequest{Solution: solution}
	if err = json.NewEncoder(conn).Encode(sr); err != nil {
		return "", fmt.Errorf("send PoW solution: %w", err)
	}

	qr := quoteResponse{}
	if err = json.NewDecoder(conn).Decode(&qr); err != nil {
		return "", fmt.Errorf("get quote: %w", err)
	}

	return qr.Quote, nil
}
