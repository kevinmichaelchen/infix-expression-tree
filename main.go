package main

import (
	"fmt"
	"github.com/kevinmichaelchen/infix-expression-tree/internal/expression"
	"strings"
)

func main() {
	in := "(20- (8-3)) - 3"
	result := compute(in)
	fmt.Println("The answer is:", result)
}

func compute(in string) int {
	expr := sanitize(in)
	tree := expression.Parse(expr)
	return tree.Resolve()
}

func sanitize(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
