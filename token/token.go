package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL="ILLEGAL"
	EOF="EOF"
	IDENT="IDENT"
	INT="INT"
	ASSIGN="="
	COMMA=","
	SEMICOLON=";"
	LPAREN="("
	RPAREN=")"
	LBRACE="{"
	RBRACE="}"
	LBRACKET="["
	RBRACKET="]"
	STRING = "STRING"
	// Operators
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"
	// Keywords
	FUNCTION = "FUNCTION"
	FOR="FOR"
	WHILE="WHILE"
	LET= "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	EQ = "=="
	NOT_EQ = "NOT_EQ"
	COLON = ":"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"for": FOR,
	"while": WHILE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
