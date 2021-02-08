package expression

import (
	"fmt"
	"strconv"
)

type Tree struct {
	Value string
	Left  *Tree
	Right *Tree
	// TODO probably unnecessary to have this field
	Parent *Tree
}

func (e Tree) isOperator() bool {
	return e.Value == "+" || e.Value == "-" || e.Value == "*" || e.Value == "/"
}

func (e Tree) getOperator() string {
	return e.Value
}

func (e Tree) Resolve() int {
	if !e.isOperator() {
		i, err := strconv.Atoi(e.Value)
		if err != nil {
			panic(fmt.Errorf("invalid operand: %s", e.Value))
		}
		return i
	}

	if e.Left == nil || e.Right == nil {
		panic("cannot solve binary expression. operator node is missing one of its children.")
	}

	return solve(e.Value, e.Left.Resolve(), e.Right.Resolve())
}

func solve(operatorStr string, left, right int) int {
	switch operatorStr {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		panic("unhandled operator type")
	}
}
