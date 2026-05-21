package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/TheoMinariniDeMelo/truth-table/parser"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
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

	l := len(props);

	var maps []map[string]bool = make([]map[string]bool, 1);

	m := make(map[string]bool);

	for _, prop := range props {
		m[prop] = true;
	}

	m[str], err = ast.Eval(m);


	if err != nil {
		fmt.Printf("Invalid operation: %s\n\r%s", str, err.Error());
		os.Exit(1);
	}

	maps[0] = m; 

	s := strings.Join(props, " | ");
	fmt.Print(s + " | " + str + "\n");

	for i := range powInt(2,l) {
		m := make(map[string]bool);
		

		for j, prop := range props {
			p := powInt(2, l - (j + 1)) 
			if i % (p + 1) == p {
				m[prop] = !maps[i][prop]
			} else {
				m[prop] = maps[i][prop]
			}
			fmt.Printf(" %s |", boolToString(m[prop]))
		}
		m[str], err = ast.Eval(m);
		fmt.Printf(" %s |\n", boolToString(m[str]))
		maps = append(maps, m);

		if err != nil {
			fmt.Printf("Invalid operation: %s\n\r%s", str, err.Error());
			os.Exit(1);
		}
	}
}

func boolToString(b bool) string{
	if b {
		return "true"
	} else {
		return "false"
	}
}
