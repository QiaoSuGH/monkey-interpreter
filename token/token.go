package token

/*
In Monkey tokens:
integers
identifiers
keywords
special characters
*/

type TokenType string

type Token struct {
	Type    TokenType
	Literal string //the literal value e.g. interger 5 value:5
}

//define types of token
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	//operators
	ASSIGN = "="
	PLUS   = "+"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
