package tokens

type Type string

const (
	Invalid Type = "invalid"
	Eof     Type = "eof"

	// Variable-related
	Identifier Type = "identifier"
	Number     Type = "number"
	String     Type = "string"

	// Comments
	CommentBegin   Type = "comment_begin"
	CommentContent Type = "comment_content"

	// Spacing and separators
	Space     Type = "space"
	Newline   Type = "newline"
	Semicolon Type = "semicolon"
	Comma     Type = "comma"
	Colon     Type = "colon"
	Question  Type = "question"

	// Blocks
	Lbrace   Type = "lbrace"
	Rbrace   Type = "rbrace"
	Lparen   Type = "lparen"
	Rparen   Type = "rparen"
	Lbracket Type = "lbracket"
	Rbracket Type = "rbracket"

	// Operators
	Assignment Type = "assignment"

	Add    Type = "add"
	Sub    Type = "sub"
	Mul    Type = "mul"
	Div    Type = "div"
	Mod    Type = "mod"
	Pow    Type = "pow"
	IntDiv Type = "int_div"

	Eq  Type = "eq"
	Neq Type = "neq"
	Gt  Type = "gt"
	Lt  Type = "lt"
	Gte Type = "gte"
	Lte Type = "lte"

	Not  Type = "not"
	And  Type = "and"
	Or   Type = "or"
	Xor  Type = "xor"
	Nor  Type = "nor"
	Nand Type = "nand"
)
