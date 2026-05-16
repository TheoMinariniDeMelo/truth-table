package parser

import (
	"errors"
	"fmt"
	"strings"
)

type Operation uint8

const (
	BINARY_AND Operation = iota
	BINARY_OR
	UNARY_NOT
	NONE
)

type AST struct {
	op    Operation
	prop string
	left  *AST
	right *AST
}

func (t *AST) Eval(props map[string]bool) (bool, error) {
	switch t.op {
	case BINARY_OR:
		{
			left, err1 := t.left.Eval(props);
			right, err2 := t.right.Eval(props);

			if err1 != nil {
				return false, err1
			}

			if err2 != nil {
				return false, err2
			}

			return left || right, nil
		}
	case BINARY_AND:
		{
			left, err1 := t.left.Eval(props)
			right, err2 := t.right.Eval(props)

			if err1 != nil {
				return false, err1
			}

			if err2 != nil {
				return false, err2
			}

			return left && right, nil
		}
	case UNARY_NOT:
		{
			left, err := t.left.Eval(props);

			if err != nil {
				return false, err;
			}

			return !left, nil;
		}
	case NONE:
		{
			return props[t.prop], nil
		}
	default:
		{
			return false, errors.New("Unidentified operation")
		}
	}
}

// debug
func (t *AST) Print(){
	var queue []*AST = make([]*AST, 1);
	queue = append(queue, t);


	for len(queue) > 0 {
		levelSize := len(queue);

		var s string;

		for range levelSize {
			node := queue[len(queue) - 1];

			switch node.op {
				case BINARY_AND: {
					s += "AND ";
				}
				case BINARY_OR: {
					s += "OR ";
				}
				case UNARY_NOT: {
					s += "NOT ";
				}
				case NONE: {
					s = node.prop + " ";
				}
				default: {
					s = "";
				}
			}
			if node.left != nil {
				queue = append(queue, node.left);
			}
			if node.right != nil {
				queue = append(queue, node.right);
			}
		}
		fmt.Printf("%s\n", s);
	}
}
