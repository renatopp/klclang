# KLC

https://github.com/skatsuta/monkey-interpreter/blob/master/parser/parser.go#L330

## TODO:

- Corrigir sistema de erros para evitar de dar panic e mostrar mais erros e mais informações
- Evaluation
  - Assignment to index
  - string ++ (number|list)
  - String interpolation
    - $<var>
    - templating? $a[2f]
  - Functions
    - call
    - scope
    - chain
    - recursion