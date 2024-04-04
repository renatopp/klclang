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

func TestOperator(t *testing.T) {
	run(t, "1 == 1", "1")
	run(t, "1 == 2", "0")
	run(t, "1 != 1", "0")
	run(t, "1 != 2", "1")
	run(t, "1 < 2", "1")
	run(t, "1 > 2", "0")
	run(t, "1 <= 2", "1")
	run(t, "1 >= 2", "0")

	run(t, "1 + 1", "2")
	run(t, "1 - 1", "0")
	run(t, "1 / 2", "0.500000")
	run(t, "3 * 2", "6")
	run(t, "2 ^ 3", "8")

	run(t, "+1", "1")
	run(t, "-1", "-1")
	run(t, "!1", "0")
	run(t, "!0", "1")

	run(t, "0 and 0", "0")
	run(t, "0 and 1", "0")
	run(t, "1 and 0", "0")
	run(t, "1 and 1", "1")

	run(t, "0 or 0", "0")
	run(t, "0 or 1", "1")
	run(t, "1 or 0", "1")
	run(t, "1 or 1", "1")

	run(t, "2km to m", "2000")
	run(t, "2000m to km", "2")
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
	run(t, `
		f() = 1
		f(1) = 2
		f(2, x) = 3
		f(x) = 4

		f() + f(1) + f(2, 3) + f(4)
	`, "10")
}

func TestAssignment(t *testing.T) {
	run(t, `a = 1; a = 2; a`, "2")
	run(t, `a = 1; a += 2; a`, "3")
	run(t, `a = 1; a -= 2; a`, "-1")
	run(t, `a = 1; a *= 2; a`, "2")
	run(t, `a = 1; a /= 2; a`, "0.500000")
	run(t, `a = 1; a ^= 2; a`, "1")
}
