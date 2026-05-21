package parser

import (
	"fmt"
)

type ParseError struct {
	expected []TokenType
	position uint16
}

func (p ParseError) Error() string {
	var s string
	for i, v := range p.expected {
		s += v.ToString()
		if(i < len(p.expected) - 1) {
			s+= ", " 
		}
	}
	return fmt.Sprintf("Unexpected token at %d, Expected: %s", p.position, s)
}

func ParseExpression(lx *Lexer) (*AST, error) {
	left, err := ParseStatment(lx)

	if err != nil {
		return nil, err
	}
	
	for lx.token.tokenType == TOK_AND || lx.token.tokenType == TOK_OR {
		t := lx.token;
		right, err := ParseStatment(lx)

		if err != nil {
			return nil, err
		}

		switch t.tokenType {
			case TOK_AND: { 
				left = &AST{op: BINARY_AND, left: left, right: right}
			}
			case TOK_OR:{
				left = &AST{op: BINARY_OR, left: left, right: right}
			}
		}
	}
	return left, nil
}

func ParseStatment(lx *Lexer) (*AST, error) {
	t, err := lx.Next()

	if err != nil {
		return nil, err
	}


	if t.tokenType == TOK_NOT {
		left, err := ParseStatment(lx)
		if err != nil {
			return nil, err
		}
		return &AST{left: left, op: UNARY_NOT}, nil
	}

	left, err := ParseTerm(lx)

	if err != nil {
		return nil, err
	}

	return left, nil
}

func ParseTerm(lx *Lexer) (*AST, error) {
	t :=  lx.token
	switch t.tokenType {
		case TOK_PROP : {
			lx.Next()
			return &AST{prop: t.lexeme, op: NONE}, nil
		}
		case TOK_LEFT_PAREN: {
			left, err := ParseExpression(lx)
			if err != nil {
				return nil, err
			}

			if lx.token.tokenType != TOK_RIGHT_PAREN {
				return nil, ParseError{position: lx.current, expected: []TokenType{TOK_RIGHT_PAREN}}
			}
			return left, nil
		}
	}
	return nil, ParseError{position: lx.current, expected: []TokenType{TOK_PROP, TOK_LEFT_PAREN}}
}
