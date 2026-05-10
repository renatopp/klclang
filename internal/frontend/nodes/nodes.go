package nodes

import "github.com/renatopp/klclang/internal/core"

type base struct {
	token *core.Token // Declaration token associated with the node
	docs  *core.Token // Documentation comment token associated with the node
}

func b(t *core.Token) base                        { return base{token: t} }
func (n *base) WithToken(v *core.Token) core.Node { n.token = v; return n }
func (n *base) GetToken() *core.Token             { return n.token }
func (n *base) GetSpan() *core.Span               { return n.token.Span }
func (n *base) WithDocs(v *core.Token) core.Node  { n.docs = v; return n }
func (n *base) GetDocs() *core.Token              { return n.docs }

// ----------------------------------------------------------------------------
// NODES
// ----------------------------------------------------------------------------

// // Module represents a Golden module, which may contain imports and declarations.
// type Module struct {
// 	base
// 	Module  *core.Module // Reference to the core module
// 	Imports []*Import    // Imported module names, from `import ...` statements
// 	Decls   []core.Node  // Declarations in the module, including functions, types, variables, etc.
// }

// func NewModule(t *core.Token, imports []*Import, decls []core.Node) *Module {
// 	return &Module{base: b(t), Imports: imports, Decls: decls}
// }

// // Import represents an import statement in a Golden module.
// type Import struct {
// 	base
// 	Value *String // Module path being imported
// }

// func NewImport(t *core.Token, value *String) *Import {
// 	return &Import{base: b(t), Value: value}
// }

// // LetExpr represents a variable declaration in Golden.
// type LetExpr struct {
// 	base
// 	AssignToken *core.Token // can be nil (ValueExpr will be too, but TypeExpr will not)
// 	Name        *ValueIdent
// 	// TypeExpr    core.Node // can be nil (but AssignToken and ValueExpr will not)
// 	ValueExpr core.Node // can be nil (AssignToken will be too, but TypeExpr will not)
// }

// func NewLetExpr(t *core.Token, assignToken *core.Token, name *ValueIdent, valueExpr core.Node) *LetExpr {
// 	return &LetExpr{base: b(t), AssignToken: assignToken, Name: name, ValueExpr: valueExpr}
// }

// // ValueIdent represents an identifier for a value in Golden.
// type ValueIdent struct {
// 	base
// 	Name string // Identifier name
// }

// func NewValueIdent(t *core.Token, name string) *ValueIdent {
// 	return &ValueIdent{base: b(t), Name: name}
// }

// // String represents a string literal in Golden.
// type String struct {
// 	base
// 	Value string // String value
// }

// func NewString(t *core.Token, value string) *String {
// 	return &String{base: b(t), Value: value}
// }

// // Integer represents an integer literal in Golden.
// type Integer struct {
// 	base
// 	Value int64 // Integer value
// }

// func NewInteger(t *core.Token, value int64) *Integer {
// 	return &Integer{base: b(t), Value: value}
// }

// // Float represents a floating-point literal in Golden.
// type Float struct {
// 	base
// 	Value float64 // Float value
// }

// func NewFloat(t *core.Token, value float64) *Float {
// 	return &Float{base: b(t), Value: value}
// }

// // Bool represents a boolean literal in Golden.
// type Bool struct {
// 	base
// 	Value bool // Boolean value
// }

// func NewBool(t *core.Token, value bool) *Bool {
// 	return &Bool{base: b(t), Value: value}
// }

// // UnaryOp represents a unary operation in Golden.
// type UnaryOp struct {
// 	base
// 	Op   string    // Operator, e.g., "-", "!"
// 	Expr core.Node // Operand expression
// }

// func NewUnaryOp(t *core.Token, op string, expr core.Node) *UnaryOp {
// 	return &UnaryOp{base: b(t), Op: op, Expr: expr}
// }

// // BinaryOp represents a binary operation in Golden.
// type BinaryOp struct {
// 	base
// 	Left  core.Node // Left operand expression
// 	Op    string    // Operator, e.g., "+", "-", "*", "/"
// 	Right core.Node // Right operand expression
// }

// func NewBinaryOp(t *core.Token, left core.Node, op string, right core.Node) *BinaryOp {
// 	return &BinaryOp{base: b(t), Left: left, Op: op, Right: right}
// }
