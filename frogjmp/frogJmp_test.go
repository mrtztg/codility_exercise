package frogjmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct{
		X, Y, D, Solution int
	} {
		{10, 85, 30, 3},
		{0, 0, 100, 0},
		{100, 100, 50, 0},
		{0, 40, 30, 2},
		{1000, 1001, 2000, 1},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.Solution, Solution(tc.X, tc.Y, tc.D))
	}
}