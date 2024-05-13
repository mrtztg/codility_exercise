package oddoccurrencesinarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		arr      []int
		solution int
	}{
		{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
		{[]int{2, 3, 8, 1000, 8, 1000, 7, 2, 3}, 7},
		{[]int{1}, 1},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.solution, Solution(testCase.arr))
	}
}
