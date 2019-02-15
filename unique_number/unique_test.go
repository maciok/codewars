package unique_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, float32(2), FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}))
	assert.Equal(t, float32(0.55), FindUniq([]float32{0, 0, 0.55, 0, 0}))
}
