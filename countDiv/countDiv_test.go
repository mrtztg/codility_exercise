package countdiv

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct{
		A, B, K, Solution int
	} {
		{1, 10, 3, 3},
		{10, 10, 5, 1},
		{11, 12, 5, 0},
		{20, 22, 1, 3},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.Solution, Solution(tc.A, tc.B, tc.K))
	}
}