package challenger

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/Tsapen/wow/internal/wow"
)

const (
	length = 10

	charset = `abcdefghijklmnopqrstuvwxyz` +
		`ABCDEFGHIJKLMNOPQRSTUVWXYZ` +
		`0123456789`
)

type challenger struct {
	solver wow.Solver
}

// New creates a new challenger.
func New(solver wow.Solver) *challenger {
	return &challenger{
		solver: solver,
	}
}

// Generate makes a random challenge for PoW.
func (*challenger) Generate() (string, error) {
	randomString := make([]byte, length)

	for i := range randomString {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		randomString[i] = charset[randomIndex.Int64()]
	}

	return string(randomString), nil
}

// Verify checks the client's PoW solution.
func (c *challenger) Verify(input, clientSolution string) error {
	serverSolution, err := c.solver.Solve(input)
	if err != nil {
		return fmt.Errorf("solve challenge \"%s\": %w", input, err)
	}

	if serverSolution != clientSolution {
		return wow.ErrWrongSolution
	}

	return nil
}
