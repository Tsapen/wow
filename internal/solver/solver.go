package solver

import (
	"crypto/sha256"
	"fmt"
)

type solver struct{}

// New creates solver.
func New() *solver {
	return &solver{}
}

// Solve calculates SHA-256 hash from input.
func (s *solver) Solve(input string) (string, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(input))
	if err != nil {
		return "", fmt.Errorf("calculate SHA-256 hash from string \"%s\": %w", input, err)
	}

	hashSum := hash.Sum(nil)

	return fmt.Sprintf("%x", hashSum), nil
}
