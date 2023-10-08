package solver_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Tsapen/wow/internal/solver"
)

func TestSolve(t *testing.T) {
	const exp = "71717682b2d3dc088fecdd646fa40519a62c5c53ebe2285c643e988ae258966a"

	got, err := solver.New().Solve("input_1234")
	assert.NoError(t, err)
	assert.Equal(t, exp, got)
}
