package parser

import (
	"fmt"
	"os"
	"slices"
)

func GetProps (lx Lexer) []string  {
	var props []string

	for !lx.IsCompleted() {
		t, err := lx.Next();

		if err != nil {
			fmt.Printf("Invalid expression: %s", err.Error());
			os.Exit(1);
		}

		if t.tokenType == TOK_PROP && !slices.Contains(props, t.lexeme) {
			props = append(props, t.lexeme);
		}
	}
	return props;
}
