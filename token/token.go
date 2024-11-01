package token

type Type uint8

const (
	ILLEGAL Type = iota
	EOF
	IDENT
	INT
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	LT
	GT
	EQ
	NOT_EQ
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

type Token struct {
	Type    Type
	Literal string
}

func New(tp Type, l string) Token {
	return Token{Type: tp, Literal: l}
}

func FromByte(tp Type, ch byte) Token {
	return Token{Type: tp, Literal: string(ch)}
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdentifier(ident string) Type {
	if tp, ok := keywords[ident]; ok {
		return tp
	}

	return IDENT
}
