# infix-expression-tree

[![Run on Repl.it](https://repl.it/badge/github/kevinmichaelchen/infix-expression-tree)](https://repl.it/github/kevinmichaelchen/infix-expression-tree)

This program accepts an infix algebraic formula, parses it,
constructs a representative binary expression tree,
and computes the result of the formula.

## Jargon
### Trees
A binary tree is a tree in which each node has at most two children (the "left child"
and the "right child").

A binary expression tree is a binary tree that can represent algebraic or boolean 
expressions that contain both unary and binary operators.

![expression-tree](https://user-images.githubusercontent.com/5129994/107164385-cd877580-697c-11eb-96e8-84070a60c1f5.png)

Leaf nodes represent operands.
Nodes with two children represent a binary operator (e.g., addition and multiplication).
Nodes with one child represent a unary operator.

### Notation
Infix notation is used for arithmetical formulae and occurs when the operator appears 
between the operands.

### Associativity
In mathematics, the associative property is a property of some binary operations, 
which means that rearranging the parentheses in an expression will not change the result.

## Challenges
### Formula Normalization (Parens Injection)
The recursion code that solves the expression tree is fairly straight-forward.
Building the tree from a string is a bit more difficult, since parsing will have
to account for order of operations in the event that parentheses are absent.

The program needs to establish an order of operations, ideally in linear time.
This means we need to scan (just once) through a string like
```
4 / 2 ^ 2 ^ 3
```
and we need to interpret this as:
```
4 / ( 2 ^ 2 ) ^ 3
4 / ( ( 2 ^ 2 ) ^ 3 )
```

The resulting tree should look like:
```
             %
           /   \
          4     ^
              /   \
             /     \
            ^       3
          /   \
         2     2 
```

The algorithm should be linear-time.

A naive solution would involve a pass to add parentheses around exponential operations, 
followed by another pass for multiplicative ops, followed by another for additive ops.  

~~Another solution might involve one pass simply to count the number of binary operations,
which will inform the height of the expression tree.~~

Take the following generalized formula:
a + b * c * d ^ e + f ^ g * h


## Other considerations
- Unary operations (e.g., ones' complement)
- Boolean logic, as opposed to algebraic arithmetic