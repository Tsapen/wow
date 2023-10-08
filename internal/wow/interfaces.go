package wow

import (
	"context"
)

// Storage is an interface for storing and retrieving quotes.
type Storage interface {
	// Quote retrieves a random quote from the storage.
	Quote(ctx context.Context) (string, error)
}

// Challenger is an interface for generating and verifying challenges.
type Challenger interface {
	// Generate returns a new challenge as a string.
	Generate() (string, error)
	// Verify verifies if the provided solution matches the challenge.
	Verify(string, string) error
}

// Solver is an interface for solving challenges.
type Solver interface {
	// Solve takes a challenge string as input and attempts to solve it.
	Solve(string) (string, error)
}
