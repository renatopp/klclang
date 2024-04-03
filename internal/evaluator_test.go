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

func TestVariables(t *testing.T) {
	run(t, "a = 1", "1")
	run(t, "a = 1; a", "1")
	run(t, "a = 1; b = 2; a + b", "3")
	run(t, "a = 1; b = 2; a + b; a = 3; a + b", "5")
}

func TestFunctions(t *testing.T) {
	run(t, "f() = 1; f()", "1")
	run(t, "f(x) = x*2; f(5)", "10")
	run(t, `
		f(0) = 1;
		f(1) = 2;
		f(x) = x*2;

		f(0) + f(1) + f(5)
	`, "13")
	run(t, `
		fib(0) = 1;
		fib(1) = 1;
		fib(n) = fib(n-1) + fib(n-2);
		fib(5)
	`, "8")
	run(t, `
		fact(0) = 1
		fact(n) = n * fact(n-1)
		fact(5)
	`, "120")
}
