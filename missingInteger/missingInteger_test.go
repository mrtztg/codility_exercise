package missinginteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct{
		A []int
		Solution int
	}{
		{[]int{1, 3, 6, 4, 1, 2}, 5},
		{[]int{1, 2, 3}, 4},
		{[]int{-1, -3}, 1},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.Solution, Solution(tc.A))
	}
}