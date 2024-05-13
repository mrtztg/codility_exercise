package binarygap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, Solution(9))
	assert.Equal(t, 4, Solution(529))
	assert.Equal(t, 1, Solution(20))
	assert.Equal(t, 0, Solution(15))
	assert.Equal(t, 0, Solution(32))
	assert.Equal(t, 5, Solution(1041))
}