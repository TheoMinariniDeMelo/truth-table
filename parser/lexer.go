package parser

import (
	"errors"
	"fmt"
	"unicode"
)

type UnidentifiedTokenError struct {
	position uint16;
};

func (u UnidentifiedTokenError) Error() string {
	return fmt.Sprintf("Unidentified token at %d", u.position);
}

type TokenType uint8;

const (
	TOK_AND TokenType = iota;
	TOK_NOT;
	TOK_OR;
	TOK_LEFT_PAREN;
	TOK_RIGHT_PAREN;
	TOK_PROP;
	TOK_NONE;
	TOK_EOF;
)

func (t TokenType) ToString() string{
	switch t {
	case TOK_AND:
		return "AND";
	case TOK_OR:
		return "OR";
	case TOK_NOT:
		return "NOT";
	case TOK_LEFT_PAREN:
		return "PAREN_LEFT";
	case TOK_RIGHT_PAREN:
		return "PAREN_RIGHT";
	case TOK_PROP:
		return "Proposition"
	default:
		return ""
	}
}

type Token struct {
	lexeme string;
	tokenType TokenType;
}

type Lexer struct {
	src string;
	current uint16;
	token Token;
}

func NewLexer(src string) *Lexer {
	return &Lexer{src, 0, Token{tokenType: TOK_NONE}};
}

func (l *Lexer) peek() byte {
	c := l.src[l.current];
	l.current++;
	return c;
}

func (l *Lexer) Next() (Token, error) {
	if len(l.src) == 0 {
		return Token{}, errors.New("the logical statment is empty");
	}
	if l.current == uint16(len(l.src)) {
		return Token{ tokenType: TOK_EOF }, nil;
	}
	l.skipWhiteSpaces();

	var tok Token;

	c := l.peek();

	if unicode.IsLetter(rune(c)) {
		tok.lexeme = string(c);
		tok.tokenType = TOK_PROP;
		l.token = tok;
		return tok, nil;
	}

	switch c {
	case '(':
		tok.lexeme = string(c);
		tok.tokenType = TOK_LEFT_PAREN;
	case ')':
		tok.lexeme = string(c);
		tok.tokenType = TOK_RIGHT_PAREN;
	case '~':
		tok.lexeme = string(c);
		tok.tokenType = TOK_NOT;
	case '^':
		tok.lexeme = string(c);
		tok.tokenType = TOK_AND;
	case '|':
		tok.lexeme = string(c);
		tok.tokenType = TOK_OR;
	default:
		return Token{}, UnidentifiedTokenError{position: l.current}; 
	};

	l.token = tok;
	return tok, nil;
}

func (l *Lexer) skipWhiteSpaces() {
	for i := l.current + 1; i < uint16(len(l.src)); i++ {
		if(l.src[l.current] != ' ') {
			return;
		}
		l.current += 1;
	}
}

func (l *Lexer) IsCompleted() bool {
	return l.current >= uint16(len(l.src));
}


func (t *Token) ToString() string {
	return t.tokenType.ToString() + ": " + t.lexeme;
}
