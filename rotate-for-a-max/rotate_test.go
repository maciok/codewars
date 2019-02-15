package rotate_kata

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, int64(85821534), MaxRot(38458215))
	assert.Equal(t, int64(988103115), MaxRot(195881031))
	assert.Equal(t, int64(962193428), MaxRot(896219342))
}