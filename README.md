# KLC

https://github.com/skatsuta/monkey-interpreter/blob/master/parser/parser.go#L330

## TODO:

- Corrigir sistema de erros para evitar de dar panic e mostrar mais erros e mais informações
- Evaluation
  - Assignment to index
  - string ++ (number|list)
  - Control flow
    - if return
    - if true false
  - String interpolation
    - $<var>
    - templating? $a[2f]
  - Functions
    - definition
    - call
    - recursion
    - scope