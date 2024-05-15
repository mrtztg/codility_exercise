package maxcounters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct{
		N int
		A []int
		Solution []int
	} {
		{5, []int{3,4,4,6,1,4,4}, []int{3,2,2,4,2}},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.Solution, Solution(tc.N, tc.A))
	}
}