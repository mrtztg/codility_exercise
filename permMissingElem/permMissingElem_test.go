package permmissingelem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct{
		A []int
		Solution int
	}{
		{[]int{2, 3, 1, 5}, 4},
		{[]int{1, 6, 8, 2, 5, 3}, 4},
		{[]int{1}, 1},
		{[]int{}, 1},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.Solution, Solution(tc.A))
	}
}