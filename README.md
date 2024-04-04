# KLC Language

KLC (pronounced `calc`) is a simple toy programming language for basic math, created for studying programming language design and development.

[Try it online](https://klc.r2p.dev/)

## Why?

KLC was created primarily for fun and learning design and development of programming languages, and secondarily for daily math - calculus that requires conversions and a bit more than what calculator apps usually provides.

## Install

```bash
curl -Ls https://github.com/renatopp/klclang/releases/download/0.1.0/klclang_Linux_x86_64.tar.gz | sudo tar -xvzf - -C /usr/local/bin
```

## Usage

You can evaluate any expression inline by using quotes `'<expression>'`, prefer single quotes because bash may interpolate double quotes: 

```bash
$ klc '1 + 2'
3

$ klc 'fn(x) = x*2; fn(4)'
8
```

Alternatively, you can execute from file:

```bash
$ klcc file.klc
1
```

If you run `klc` without any arguments, you access the REPL, an interactive shell where you can test your commands.

## How does it looks?

```haskell
-- Comments are like these and can be used as documentation.
my_constant = 10
help(my_constant) -- will print 'Comments are like these and can be used as documentation.'

-- You can create variables
a = 1
b = 2
c = a + b

-- You can create functions using pattern matching
fib(0) = 0
fib(1) = 1
fib(x) = fib(x-1) + fib(x-2)

-- You can use the metrics system
distance = 2.5km to m -- converts 2.5 kilometers to meters
size = 40321MB to GB -- convert 40321 megabytes to gigabytes

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
!1      -- not
0 and 1
0 or 1
```

## Documentation

- [Language Details](./.docs/language.md)
- [Builtin Functions and Constants](./.docs/builtin.md)

## More Info

- [https://r2p.dev](https://r2p.dev)