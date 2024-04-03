package internal_test

import (
	"testing"

	"github.com/renatopp/klclang/internal"
	"github.com/stretchr/testify/assert"
)

func run(t *testing.T, code string, expected string) {
	t.Helper()
	obj, err := internal.Run([]byte(code))
	assert.NotNil(t, obj)
	assert.Nil(t, err)
	if obj != nil {
		assert.Equal(t, expected, obj.String())
	}
}

func TestExpressions(t *testing.T) {
	run(t, "1", "1")
	run(t, "1.2", "1.200000")
	run(t, "0", "0")
	run(t, "+1", "1")
	run(t, "-1", "-1")
	run(t, "1 + 2", "3")
	run(t, "1 + 2 * 3", "7")
	run(t, "(1 + 2) * 3", "9")
	run(t, "2 * -3", "-6")
	run(t, "2 * -3 + 4", "-2")
	run(t, "2^-2", "0.250000")
	run(t, "-2^2", "4")
}
