package sum_kata

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, [2]int{0, 2}, TwoSum([]int{1, 2, 3}, 4))
	assert.Equal(t, [2]int{1, 2}, TwoSum([]int{1234, 5678, 9012}, 14690))
	assert.Equal(t, [2]int{0, 1}, TwoSum([]int{2, 2, 3}, 4))
}
