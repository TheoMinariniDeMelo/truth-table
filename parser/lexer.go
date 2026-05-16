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
)

func (t TokenType) ToString() string{
  switch t {
  case TOK_AND:
    return "^";
  case TOK_OR:
    return "|";
  case TOK_NOT:
    return "~";
  case TOK_LEFT_PAREN:
    return "(";
  case TOK_RIGHT_PAREN:
    return ")";
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
}

func NewLexer(src string) Lexer {
  return Lexer{src, 0};
}

func (l *Lexer) peek() byte {
  return l.src[l.current];
}

func (l *Lexer) Next() (Token, error) {
  if len(l.src) == 0 {
    return Token{}, errors.New("the logical statment is empty");
  }

  l.skipWhiteSpaces();

  var tok Token;

  c := l.peek();

  if unicode.IsLetter(rune(c)) {
    tok.lexeme = string(c);
    tok.tokenType = TOK_PROP;
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

  return tok, nil;
}

func (l *Lexer) skipWhiteSpaces() {
  for i := l.current + 1; i < uint16(len(l.src)); i++ {
    if(l.src[l.current] != ' ') {
      return;
    }
  }
}
