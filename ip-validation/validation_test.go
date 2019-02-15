package ip_validation

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, true, Is_valid_ip("12.255.56.1"))
	assert.Equal(t, false, Is_valid_ip(""))
	assert.Equal(t, false, Is_valid_ip("abc.def.ghi.jkl"))
	assert.Equal(t, false, Is_valid_ip("123.456.789.0"))
	assert.Equal(t, false, Is_valid_ip("12.34.56"))
	assert.Equal(t, false, Is_valid_ip("12.34.56 .1"))
	assert.Equal(t, false, Is_valid_ip("123.045.067.089"))
	assert.Equal(t, true, Is_valid_ip("127.1.1.0"))
	assert.Equal(t, true, Is_valid_ip("0.0.0.0"))
	assert.Equal(t, true, Is_valid_ip("0.34.82.53"))
	assert.Equal(t, false, Is_valid_ip("192.168.1.300"))
}
