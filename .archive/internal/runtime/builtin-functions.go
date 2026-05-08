package runtime

import (
	"math"
	"math/rand/v2"
)

// MATH
var (
	FN_ABS = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Abs(n))
	}), "Abs returns the absolute value of a number.")

	FN_ACOS = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Acos(n))
	}), "Acos returns the arccosine of a number.")

	FN_ACOSH = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Acosh(n))
	}), "Acosh returns the hyperbolic arccosine of a number.")

	FN_ASIN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Asin(n))
	}), "Asin returns the arcsine of a number.")

	FN_ASINH = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Asinh(n))
	}), "Asinh returns the hyperbolic arcsine of a number.")

	FN_ATAN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Atan(n))
	}), "Atan returns the arctangent of a number.")

	FN_ATAN2 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		y := args[0].Number()
		x := args[1].Number()
		return NewNumber(math.Atan2(y, x))
	}), "Atan2 returns the arctangent of y/x, using the signs of the two to determine the quadrant of the result.")

	FN_ATANH = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Atanh(n))
	}), "Atanh returns the hyperbolic arctangent of a number.")

	FN_CBRT = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Cbrt(n))
	}), "Cbrt returns the cube root of a number.")

	FN_CEIL = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Ceil(n))
	}), "Ceil returns the smallest integer value greater than or equal to a number.")

	FN_COPYSIGN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		return NewNumber(math.Copysign(x, y))
	}), "Copysign returns the first argument with the sign of the second argument.")

	FN_COS = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Cos(n))
	}), "Cos returns the cosine of a number.")

	FN_COSH = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(math.Cosh(n))
	}), "Cosh returns the hyperbolic cosine of a number.")

	FN_DIM = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		return NewNumber(math.Dim(x, y))
	}), "Dim returns the maximum of x-y or 0.")

	FN_ERF = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Erf(x))
	}), "Erf returns the error function of a number.")

	FN_ERFC = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Erfc(x))
	}), "Erfc returns the complementary error function of a number.")

	FN_ERFCINV = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Erfcinv(x))
	}), "Erfcinv returns the inverse of the complementary error function of a number.")

	FN_ERFINV = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Erfinv(x))
	}), "Erfinv returns the inverse of the error function of a number.")

	FN_EXP = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Exp(x))
	}), "Exp returns e**x, the base-e exponential of x.")

	FN_EXP2 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Exp2(x))
	}), "Exp2 returns 2**x, the base-2 exponential of x.")

	FN_EXPM1 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Expm1(x))
	}), "Expm1 returns e**x - 1, the base-e exponential of x minus 1.")

	FN_FMA = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 3)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		z := args[2].Number()
		return NewNumber(math.FMA(x, y, z))
	}), "FMA returns x*y+z with no intermediate rounding.")

	FN_FLOOR = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Floor(x))
	}), "Floor returns the largest integer value less than or equal to a number.")

	FN_GAMMA = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Gamma(x))
	}), "Gamma returns the gamma function of a number.")

	FN_HYPOT = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		return NewNumber(math.Hypot(x, y))
	}), "Hypot returns the square root of x*x + y*y.")

	FN_ILOGB = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(float64(math.Ilogb(x)))
	}), "Ilogb returns the binary exponent of a number.")

	FN_INF = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 0)); err != nil {
			return err
		}

		return NewNumber(math.Inf(1))
	}), "Inf returns positive infinity.")

	FN_ISINF = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		if math.IsInf(x, 1) {
			return NewNumber(1)
		}
		return NewNumber(0)
	}), "IsInf reports whether a number is positive infinity.")

	FN_J0 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.J0(x))
	}), "J0 returns the Bessel function of the first kind of order 0.")

	FN_J1 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.J1(x))
	}), "J1 returns the Bessel function of the first kind of order 1.")

	FN_JN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		n := int(args[0].Number())
		x := args[1].Number()
		return NewNumber(math.Jn(n, x))
	}), "Jn returns the Bessel function of the first kind of order n.")

	FN_LDEXP = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		exp := int(args[1].Number())
		return NewNumber(math.Ldexp(x, exp))
	}), "Ldexp returns x * 2**exp.")

	FN_LGAMMA = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		lg, _ := math.Lgamma(x)
		return NewNumber(lg)
	}), "Lgamma returns the natural logarithm of the absolute value of the gamma function of a number.")

	FN_LOG = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Log(x))
	}), "Log returns the natural logarithm of a number.")

	FN_LOG10 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Log10(x))
	}), "Log10 returns the decimal logarithm of a number.")

	FN_LOG1P = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Log1p(x))
	}), "Log1p returns the natural logarithm of 1 plus a number.")

	FN_LOG2 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Log2(x))
	}), "Log2 returns the binary logarithm of a number.")

	FN_MAX = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		return NewNumber(math.Max(x, y))
	}), "Max returns the larger of two numbers.")

	FN_MIN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		return NewNumber(math.Min(x, y))
	}), "Min returns the smaller of two numbers.")

	FN_MOD = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		return NewNumber(math.Mod(x, y))
	}), "Mod returns the floating-point remainder of x/y.")

	FN_INT = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		intPart, _ := math.Modf(x)
		return NewNumber(intPart)
	}), "Int returns the integer parts of a number.")

	FN_FRAC = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		_, fractPart := math.Modf(x)
		return NewNumber(fractPart)
	}), "Frac returns the fractional parts of a number.")

	FN_NEXTAFTER = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		return NewNumber(math.Nextafter(x, y))
	}), "Nextafter returns the next representable float64 value after x in the direction of y.")

	FN_POW = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		return NewNumber(math.Pow(x, y))
	}), "Pow returns x**y, the base-x exponential of y.")

	FN_POW10 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		n := int(args[0].Number())
		return NewNumber(math.Pow10(n))
	}), "Pow10 returns 10**n, the base-10 exponential of n.")

	FN_REMAINDER = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		x := args[0].Number()
		y := args[1].Number()
		return NewNumber(math.Remainder(x, y))
	}), "Remainder returns the IEEE 754 floating-point remainder of x/y.")

	FN_ROUND = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Round(x))
	}), "Round returns the nearest integer, rounding half away from zero.")

	FN_ROUNDTOEVEN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.RoundToEven(x))
	}), "RoundToEven returns the nearest integer, rounding ties to even.")

	FN_SIGN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		if x > 0 {
			return NewNumber(1)
		}
		if x < 0 {
			return NewNumber(-1)
		}
		return NewNumber(0)
	}), "Sign returns the sign of a number.")

	FN_SIN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Sin(x))
	}), "Sin returns the sine of a number.")

	FN_SINH = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Sinh(x))
	}), "Sinh returns the hyperbolic sine of a number.")

	FN_SQRT = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Sqrt(x))
	}), "Sqrt returns the square root of a number.")

	FN_TAN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Tan(x))
	}), "Tan returns the tangent of a number.")

	FN_TANH = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Tanh(x))
	}), "Tanh returns the hyperbolic tangent of a number.")

	FN_TRUNC = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Trunc(x))
	}), "Trunc returns the integer value of a number.")

	FN_Y0 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Y0(x))
	}), "Y0 returns the Bessel function of the second kind of order 0.")

	FN_Y1 = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		x := args[0].Number()
		return NewNumber(math.Y1(x))
	}), "Y1 returns the Bessel function of the second kind of order 1.")

	FN_YN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		n := int(args[0].Number())
		x := args[1].Number()
		return NewNumber(math.Yn(n, x))
	}), "Yn returns the Bessel function of the second kind of order n.")
)

// RANDOM
var (
	FN_RAND = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 0)); err != nil {
			return err
		}

		return NewNumber(rand.Float64())
	}), "Rand returns a random number in the range [0.0, 1.0).")

	FN_RANDEXP = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		lambda := args[0].Number()
		return NewNumber(rand.ExpFloat64() / lambda)
	}), "RandExp returns a random number drawn from an exponential distribution with rate lambda.")

	FN_RANDINT = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 2)); err != nil {
			return err
		}

		min := int(args[0].Number())
		max := int(args[1].Number())
		return NewNumber(float64(rand.IntN(max-min) + min))
	}), "RandInt returns a random integer in the range [min, max).")

	FN_RANDINTN = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 0)); err != nil {
			return err
		}

		n := args[0].Number()
		return NewNumber(float64(rand.IntN(int(n))))
	}), "RandIntn returns a random integer in the range [0, max).")

	FN_RANDNORM = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 0)); err != nil {
			return err
		}

		return NewNumber(rand.NormFloat64())
	}), "RandNorm returns a random number drawn from a normal distribution with mean 0 and standard deviation 1.")
)

// OTHER
var (
	FN_ASSERT = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		if args[0].Number() != 0 {
			return NewNumber(1)
		}

		return NewError("assertion failed")
	}), "Assert returns true if the argument is true, otherwise it returns an error message.")

	FN_HELP = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		if err := withChecks(nArgs(args, 1)); err != nil {
			return err
		}

		println(args[0].Docs())
		return NewNumber(0)
	}), "Help prints the documentation of a function.")
)

func registerFunctions(scope *Scope) {
	scope.Set("abs", FN_ABS)
	scope.Set("acos", FN_ACOS)
	scope.Set("acosh", FN_ACOSH)
	scope.Set("asin", FN_ASIN)
	scope.Set("asinh", FN_ASINH)
	scope.Set("atan", FN_ATAN)
	scope.Set("atan2", FN_ATAN2)
	scope.Set("atanh", FN_ATANH)
	scope.Set("cbrt", FN_CBRT)
	scope.Set("ceil", FN_CEIL)
	scope.Set("copySign", FN_COPYSIGN)
	scope.Set("cos", FN_COS)
	scope.Set("cosh", FN_COSH)
	scope.Set("dim", FN_DIM)
	scope.Set("erf", FN_ERF)
	scope.Set("erfc", FN_ERFC)
	scope.Set("erfcinv", FN_ERFCINV)
	scope.Set("erfinv", FN_ERFINV)
	scope.Set("exp", FN_EXP)
	scope.Set("exp2", FN_EXP2)
	scope.Set("expm1", FN_EXPM1)
	scope.Set("fma", FN_FMA)
	scope.Set("floor", FN_FLOOR)
	scope.Set("gamma", FN_GAMMA)
	scope.Set("hypot", FN_HYPOT)
	scope.Set("ilogb", FN_ILOGB)
	scope.Set("inf", FN_INF)
	scope.Set("isInf", FN_ISINF)
	scope.Set("j0", FN_J0)
	scope.Set("j1", FN_J1)
	scope.Set("jn", FN_JN)
	scope.Set("ldexp", FN_LDEXP)
	scope.Set("lgamma", FN_LGAMMA)
	scope.Set("log", FN_LOG)
	scope.Set("log10", FN_LOG10)
	scope.Set("log1p", FN_LOG1P)
	scope.Set("log2", FN_LOG2)
	scope.Set("max", FN_MAX)
	scope.Set("min", FN_MIN)
	scope.Set("mod", FN_MOD)
	scope.Set("int", FN_INT)
	scope.Set("frac", FN_FRAC)
	scope.Set("nextAfter", FN_NEXTAFTER)
	scope.Set("pow", FN_POW)
	scope.Set("pow10", FN_POW10)
	scope.Set("remainder", FN_REMAINDER)
	scope.Set("round", FN_ROUND)
	scope.Set("roundToEven", FN_ROUNDTOEVEN)
	scope.Set("sign", FN_SIGN)
	scope.Set("sin", FN_SIN)
	scope.Set("sinh", FN_SINH)
	scope.Set("sqrt", FN_SQRT)
	scope.Set("tan", FN_TAN)
	scope.Set("tanh", FN_TANH)
	scope.Set("trunc", FN_TRUNC)
	scope.Set("y0", FN_Y0)
	scope.Set("y1", FN_Y1)
	scope.Set("yn", FN_YN)

	scope.Set("rand", FN_RAND)
	scope.Set("randExp", FN_RANDEXP)
	scope.Set("randInt", FN_RANDINT)
	scope.Set("randIntn", FN_RANDINTN)
	scope.Set("randNorm", FN_RANDNORM)

	scope.Set("assert", FN_ASSERT)
	scope.Set("help", FN_HELP)
}

func nArgs(args []Object, n int) Object {
	if len(args) != n {
		return NewError("expected %d arguments, got %d", n, len(args))
	}

	return nil
}

func withChecks(c ...Object) Object {
	for _, o := range c {
		if o != nil {
			return o
		}
	}

	return nil
}
