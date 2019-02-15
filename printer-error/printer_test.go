package print_kata

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "3/56", PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
	assert.Equal(t, "6/60", PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
	assert.Equal(t, "11/65", PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu"))
}
