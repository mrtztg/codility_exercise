package frogriverone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		X        int
		A        []int
		Solution int
	}{
		{5, []int{1, 3, 1, 4, 2, 3, 5, 4}, 6},
		{1, []int{1, 1, 1, 1}, 0},
		{4, []int{}, -1},
		{5, []int{1, 3, 5}, -1},
		{1, []int{1}, 0},
		{2, []int{2, 1}, 1},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.Solution, Solution(tc.X, tc.A))
	}
}