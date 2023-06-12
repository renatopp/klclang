# KLC Language

An experimental non-finished language that I used to started learning how to create a new programming language. Do not use to launch a rocket.

## Types:

There is only numbers, lists and strings:

```
# Numbers are the same type:
a = 1
b = 0.1234
c = -10

# Strings are single quoted:
a = 'Hello, World'

# Lists are made of other stuff:
a = [1, 2, 3, 4]
b = [0, a, 'string']
```

## Metric System

KLC provides a builtin metrics constants that are useful for conversion:

```
> 1m # one meter
1

> 1km # one kilometer
1000

> 1cm to km # syntax sugar for (1*cm)/km
0.000010

> 2GB to MB
2048

> 32h to s
115200
```

## Functions

```
fib = fn x {
  ? x <= 0 : 0 # syntax sugar for 'if x<0 { return 0 }'
  ? x == 1 : 1
  fib(x - 1) + fib(x - 2)
}

assert fib(10) == 144 # parenthesis are optional when calling functions with arguments
```

You can use pipe-like functions where the result of the previous expression is inserted as the first argument of the next function:

```
euler1 = fn(n=1000) {
  range(n)
    .filter(where x%3 == 0 || x%5 == 0)
    .sum()
}

assert euler1(3) == 0
assert euler1(5) == 3
assert euler1(6) == 8
```

Notice that `where <expression>` is a syntax sugar for `fn x=0, y=0, z=0 { <expression> }`.

Another useful syntax sugar is the `is` keyword:

```
> range(10).filter(where x is even)
[0, 2, 4, 6, 8]
```

## Control flow

There is no loops, use recursion OR `range()` and `filter`, `map`, `reduce`, `each` instead.

Ifs are only ternaries:

```
2 is even ? 'Yes' : 'No'
```

## Other Examples

### Quicksort

```
quicksort = fn ...list {
  ? !len(list) : list
  p = list[len(list)//2]
	
  quicksort(filter list where x < p) ++
           (filter list where x == p) ++
  quicksort(filter list where x > p)
}

quicksort([5, 7, 2, 1, 0, 3, 8, 3, 2])
```

### Euler #2

```
solution = fn max, a=1, b=1, sum=0 {
  c = a + b
  ? c > max: sum
  ? solution max, b, c, c is even ? sum + c : sum
}

assert solution(10) == 10
assert solution(50) == 44

solution 4000000
``` 


### Euler #3

```
solution = fn n, i=2 {
  ? i*i > n : n
  ? n % i : solution(n, i+1)
  ? solution(n//i, i)
}

solution(600851475143)
```
