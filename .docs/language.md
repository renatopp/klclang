# The Language

## Numbers and Variables

In KLC, all numbers are treated as `float64` data type. This means you can use both integers and floating-point numbers in your calculations. Here's how you can declare numbers:

```
n1 = 100
n2 = 100.0
n3 = 10e2
```

Variables in KLC are declared using the = operator. The variable name is on the left, and the value you want to assign to the variable is on the right. Once a variable is declared, you can use it in your calculations. For example:

```
a = 10
b = 20
sum = a + b -- 30
```

In the above example, sum will hold the value 300, which is the result of adding n1 and n2.

Variables must start with a letter and may contain letters, digits and `_` character.

Comments are created using the `--` symbol. Anything after `--` on a line is ignored by the interpreter. This allows you to add notes or describe what your code is doing without affecting the program execution. Additionally, comments on the same line (or right above it, instead) as a variable declaration are often used to provide documentation or help for that variable. This can be particularly useful for explaining the purpose or usage of a variable in a larger program.

```
-- Hello, World!
foo = 1

-- This comment will be supressed
bar = 2 -- by me!

help(foo)
help(bar)
```

Resulting in:

```base
Hello, World!
by me!
```

## Operations

KLC supports a variety of operators for performing calculations and comparisons.

```
-- Arithmetic Operators:
5+10 -- add
5-10 -- sub
5*10 -- mul
5/10 -- div
5^10 -- pow
5%10 -- mod

-- Relational Operators:
5 == 10 -- eq
5 != 10 -- neq
5 < 10  -- lt
5 <= 10 -- lte
5 > 10  -- gt
5 >= 10 -- gte

-- Logical Operators:
!1		   -- not
0 and 1
0 or 1
```

Booleans are represented as numbers, where `0` is false and any non-zero value is true.

All arithmetic operators can also be used in the assignment shortcuts:

```
a = 1
a += 1
a -= 1
a /= 1
a *= 1
a ^= 1
a %= 1
```

### Metric System

KLC supports the metric system, allowing you to write expressions using metric units. This feature makes it easier to perform calculations involving units of measurement.

You can write expressions like `10H` and `10M to S`. Here, `10H` is equivalent to `10*H`, where `H` is a constant that represents hours.

You can also convert between units using the `to` keyword. For example, `10H to S` is equivalent to `(10*H)/S`, converting hours to seconds.

KLC includes many metric constants for different units of measurement. Each class of metric (time, length, weight, etc.) uses a default reference value, which are all equal to 1:

- Time: the reference unit is the second (`S`).
- Length: the reference unit is the meter (`m`).
- Weight: the reference unit is the kilogram (`kg`).
- Angle: the reference unit is the radian (`rad`).
- Area: the reference unit is the square meter (`sqm`).
- Bytes: the reference unit is the byte (`B`).
- Volume: the reference unit is the liter (`l`).

This means that when you write `10H`, KLC interprets it as `10` times the number of seconds in an hour. Similarly, `10M to S` is interpreted as `10` minutes converted to the equivalent number of seconds.

## Functions

Functions are predefined procedures that perform a specific task. They take in parameters, perform an operation, and often return a result. You can call a function by writing its name followed by parentheses `()`. Inside the parentheses, you provide the input values (or arguments) for the function.

For example, the `floor(value)` function in KLC takes a number as an input and returns the largest integer less than or equal to that number.

Here's how you can call the `floor` function:

```
result = floor(3.14) -- result will hold the value 3
```

### Custom Functions

Pattern matching allows us to define multiple versions of a function, each with different parameters. The correct version of the function is then chosen based on the parameters provided when the function is called.

For example, the `factorial` function is defined with two versions:

```
factorial(0) = 1
factorial(x) = x * factorial(-1)
```
The first version is for the base case of the factorial function, where the factorial of 0 is defined as 1. The second version calculates the factorial of a number x by multiplying x with the factorial of x-1.

Another example. The `fib` function is also defined with three versions:

```
fib(0) = 0
fib(1) = 1
fib(x) = fib(x-1) + fib(x-2)
```
The first two versions define the base cases of the Fibonacci sequence, where the first and second terms are 0 and 1, respectively. The third version calculates the xth term of the Fibonacci sequence by adding the (x-1)th and (x-2)th terms.

```
f() = 1
f(1) = 2
f(2, x) = 3
f(x) = 4

assert( f() + f(1) + f(2, 3) + f(4) == 10 )
```