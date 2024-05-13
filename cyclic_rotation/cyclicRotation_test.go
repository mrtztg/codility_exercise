package cyclicrotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []int{6, 3, 8, 9, 7}, Solution([]int{3, 8, 9, 7, 6}, 1))
	assert.Equal(t, []int{7, 6, 3, 8, 9}, Solution([]int{3, 8, 9, 7, 6}, 2))
	assert.Equal(t, []int{9, 7, 6, 3, 8}, Solution([]int{3, 8, 9, 7, 6}, 3))
	assert.Equal(t, []int{0, 0, 0}, Solution([]int{0, 0, 0}, 1))
	assert.Equal(t, []int{1, 2, 3, 4}, Solution([]int{1, 2, 3, 4}, 4))
}
