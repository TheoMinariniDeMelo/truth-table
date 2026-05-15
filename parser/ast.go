package parser

import "errors"

type Operation uint8

const (
	BINARY_AND Operation = iota
	BINARY_OR
	UNARY_NOT
	NONE
)

type AST struct {
	op    Operation
	value bool
	left  *AST
	right *AST
}

func (t *AST) Eval() (bool, error) {
	switch t.op {
	case BINARY_OR:
		{
			left, err1 := t.left.Eval()
			right, err2 := t.right.Eval()

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
			left, err1 := t.left.Eval()
			right, err2 := t.right.Eval()

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
			return !t.value, nil
		}
	case NONE:
		{
			return t.value, nil
		}
	default:
		{
			return false, errors.New("Unidentified operation")
		}
	}
}
