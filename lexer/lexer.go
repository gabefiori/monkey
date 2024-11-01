package lexer

import (
	"bytes"
	"io"
	"monkey/token"
)

type Lexer struct {
	readerBuffer []byte
	reader       io.Reader

	identBuffer *bytes.Buffer

	ch     byte
	peekCh byte
}

func New(r io.Reader) *Lexer {
	l := &Lexer{
		reader:       r,
		readerBuffer: make([]byte, 1),
		identBuffer:  bytes.NewBuffer([]byte{}),
	}

	l.readChar()

	return l
}

func (l *Lexer) Prepare() {
	l.readChar()
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	var tok token.Token

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = token.New(token.EQ, string([]byte{l.ch, l.peekCh}))
			l.readChar()
		} else {
			tok = token.FromByte(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.FromByte(token.PLUS, l.ch)
	case '-':
		tok = token.FromByte(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok = token.New(token.NOT_EQ, string([]byte{l.ch, l.peekCh}))
			l.readChar()
		} else {
			tok = token.FromByte(token.BANG, l.ch)
		}
	case '/':
		tok = token.FromByte(token.SLASH, l.ch)
	case '*':
		tok = token.FromByte(token.ASTERISK, l.ch)
	case '<':
		tok = token.FromByte(token.LT, l.ch)
	case '>':
		tok = token.FromByte(token.GT, l.ch)
	case ';':
		tok = token.FromByte(token.SEMICOLON, l.ch)
	case ',':
		tok = token.FromByte(token.COMMA, l.ch)
	case '{':
		tok = token.FromByte(token.LBRACE, l.ch)
	case '}':
		tok = token.FromByte(token.RBRACE, l.ch)
	case '(':
		tok = token.FromByte(token.LPAREN, l.ch)
	case ')':
		tok = token.FromByte(token.RPAREN, l.ch)
	case 0:
		tok = token.New(token.EOF, "")
	default:
		if isLetter(l.ch) {
			literal := l.readLiteral(isLetter)

			return token.New(token.LookupIdentifier(literal), literal)
		}

		if isDigit(l.ch) {
			return token.New(token.INT, l.readLiteral(isDigit))
		}

		return token.FromByte(token.ILLEGAL, l.ch)
	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.peekCh != 0 {
		l.ch = l.peekCh
		l.peekCh = 0

		return
	}

	_, err := l.reader.Read(l.readerBuffer)

	if err != nil {
		if err == io.EOF {
			l.ch = 0
		} else {
			panic(err)
		}
	} else {
		l.ch = l.readerBuffer[0]
	}
}

func (l *Lexer) peekChar() byte {
	if l.peekCh != 0 {
		return l.peekCh
	}

	_, err := l.reader.Read(l.readerBuffer)

	if err != nil {
		if err == io.EOF {
			l.peekCh = 0
		} else {
			panic(err)
		}
	} else {
		l.peekCh = l.readerBuffer[0]
	}

	return l.peekCh
}

func (l *Lexer) readLiteral(checkFn func(byte) bool) string {
	l.identBuffer.Reset()

	for checkFn(l.ch) {
		l.identBuffer.WriteByte(l.ch)
		l.readChar()
	}

	return l.identBuffer.String()
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
