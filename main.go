package main

import (
	"fmt"
	"os"
	"github.com/TheoMinariniDeMelo/truth-table/parser"
)

func main(){
	str := "~p^q|(p^~q|k)";

	var lx *parser.Lexer = parser.NewLexer(str);

//	for t, err := lx.Next(); !lx.IsCompleted(); t, err = lx.Next() {
//		if lxError, ok := err.(parser.UnidentifiedTokenError); ok {
//			fmt.Printf("Invalid expression: %s\n\r%s", str, lxError.Error());
//			os.Exit(1);
//		}
//		fmt.Printf("%s, ", t.ToString());
//	}
	ast, err := parser.ParseExpression(lx);
	if lxError, ok := err.(parser.UnidentifiedTokenError); ok {
		fmt.Printf("Invalid expression: %s\n\r%s", str, lxError.Error());
		os.Exit(1);
	}
	if lxError, ok := err.(parser.ParseError); ok {
		fmt.Printf("Invalid operation: %s\n\r%s", str, lxError.Error());
		os.Exit(1);
	}
	ast.Print();
}
