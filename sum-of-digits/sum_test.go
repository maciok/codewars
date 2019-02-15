package sum_of_digits

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 7, DigitalRoot(16))
	assert.Equal(t, 6, DigitalRoot(195))
	assert.Equal(t, 2, DigitalRoot(992))
	assert.Equal(t, 9, DigitalRoot(167346))
	assert.Equal(t, 0, DigitalRoot(0))
}
