# infix-expression-tree

This program accepts an infix algebraic formula, parses it,
constructs a representative binary expression tree,
and computes the result of the formula.

## Primer on trees
A binary tree is a tree in which each node has at most two children (the "left child" and the "right child").

A binary expression tree is a binary tree that can represent algebraic or boolean expressions that contain both unary and binary operators.

Leaf nodes represent operands.
Nodes with two children represent a binary operator (e.g., addition and multiplication).
Nodes with one child represent a unary operator.

## Primer on notation
Infix notation is used for arithmetical formulae and occurs when the operator appears between the operands.