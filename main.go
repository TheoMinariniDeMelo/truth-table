package main

import (
	"fmt"
	"os"
	"github.com/TheoMinariniDeMelo/truth-table/parser"
)

func main(){
	str := "~p^q|(p^~q|k)";

	var lx *parser.Lexer = parser.NewLexer(str);
	
	var props []string = parser.GetProps(*lx);

	ast, err := parser.ParseExpression(lx);

	if lxError, ok := err.(parser.UnidentifiedTokenError); ok {
		fmt.Printf("Invalid expression: %s\n\r%s", str, lxError.Error());
		os.Exit(1);
	}
	if lxError, ok := err.(parser.ParseError); ok {
		fmt.Printf("Invalid operation: %s\n\r%s", str, lxError.Error());
		os.Exit(1);
	}

	result, err := ast.Eval(props);

	if err != nil {
		fmt.Printf("Invalid operation: %s\n\r%s", str, err.Error());
		os.Exit(1);
	}
	
	var s string;

	if result {
		s = "True"
	} else {
		s = "False"
	}

	fmt.Printf("\n\nresult: %s", s);

}
