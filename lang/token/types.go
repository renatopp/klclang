package token

type Type string

const (
	Invalid Type = "invalid"
	Eof     Type = "eof"

	// Variable-related
	Identifier Type = "identifier" // [a-zA-Z_][a-zA-Z0-9_]*
	Number     Type = "number"     // 123, 123.456, 123e456, -.2f
	String     Type = "string"     // '...'

	// Comments
	Comment Type = "comment" // "#"

	// Spacing and separators
	Newline   Type = "newline"   // "\n"
	Semicolon Type = "semicolon" // ";"
	Comma     Type = "comma"     // ","
	Colon     Type = "colon"     // ":"
	Question  Type = "question"  // "?"
	Dot       Type = "dot"       // "."
	Arrow     Type = "arrow"     // "=>"
	Spread    Type = "spread"    // "..."
	Backslash Type = "backslash" // "\"

	// Blocks
	Lbrace   Type = "lbrace"   // "{"
	Rbrace   Type = "rbrace"   // "}"
	Lparen   Type = "lparen"   // "("
	Rparen   Type = "rparen"   // ")"
	Lbracket Type = "lbracket" // "["
	Rbracket Type = "rbracket" // "]"

	// Operators
	Assignment Type = "assignment" // "="
	Operator   Type = "operator"   // "+", "-", "*", "/", "%", "**", "//", "==", "!=", ">", "<", ">=", "<=", "!", "&&", "||", "^^", "!|", "!&", "++",
)
