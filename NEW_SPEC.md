


# Features

```python
# AUTO DOCS

# This is a comment that will be used as auto doc for the expression
# below. Comments can be multi-lined
e = 3
help e


# COMMANDS
help <name>
exit <code?>
clear
assert <expression>
list <prefix>
print 'renato %s %t' value value

# TYPES (only numbers)
1
1e10
-1e-10
1_000_000

# UNITS AND MEASURESunits
1m           # 1 meter
1km          # 1 kilometer
1km == 1000m # true
1km to m     # 1000 meter
1m to s      # invalid metric system length to time

metric length(m=1, ms=0.1)


# VARIABLES AND FUNCTIONS
pi = 3.14
add(a, b) = a + b
ease_in_back(x) =
  c1 = 1.70158
  c3 = c1 + 1
  c3*x*x*x - c1*x*x # return the last value


# CONTROL
ease_in_elastic(x) = 
  c4 = 2*pi/3
  {
    x == 0    ? 0
    x == 1    ? 1
    otherwise ? -2^(10*x - 10) * sin((x*10 - 10.75) * c4)
  }

# OPERATORS
5+10    # add
5-10    # sub
5*10    # mul
5/10    # div
5^10    # pow
5%10    # mod
5 == 10 # eq
5 != 10 # neq
5 < 10  # lt
5 <= 10 # lte
5 > 10  # gt
5 >= 10 # gte
!1      # not
0 and 1
0 or 1
0 xor 1

```

## Future Ideas

- Matrices (`[1 2 3]` as 1x3 matrix, `[1 2 3; 1 2 3]` as 2x3 matrix).
- Numpy-like functions (`cos([1 2 3])` is the same as `[cos(1) cos(2) cos(3)]`).
- Scalar and matrix operations (`[1 2 3]*10` is `[10 20 30]`).
- Function types (`map([1 2 3], plus_one)`).
- Function application (`fn(a, b)` could be written as `fn a b`). This includes commands.




