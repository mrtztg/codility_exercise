package tapeequilibrium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		arr []int
		solution int
	}{
		{[]int{3, 1, 2, 4, 3}, 1},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.solution, Solution(tc.arr))
	}
}
// both sides should be non-empty.
// items can be nagative.
// P will be >=1 and < len-1 or len ?

// [3,1,2,4,3] -> 13
// [3, 4, 6, 10, 13]
// p=1 -> [3][1,2,4,3] -> 3 - 10
// p=2 -> [3,1]{2,4,3} => 4 - 9
// p=3 -> [3,1,2][4,3] => 6 - 7 
