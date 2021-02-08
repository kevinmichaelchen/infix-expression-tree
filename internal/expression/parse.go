package expression

// parses an infix formula and constructs a representative binary expression tree
// e.g., "(20-(8-3))-3"
func Parse(s string) *Tree {
	root := &Tree{
		Value: "+",
	}

	l := &Tree{
		Value:  "3",
		Parent: root,
	}
	r := &Tree{
		Value:  "*",
		Parent: root,
	}

	root.Left = l
	root.Right = r

	rl := &Tree{
		Value:  "+",
		Parent: r,
	}

	rr := &Tree{
		Value:  "2",
		Parent: r,
	}
	r.Left = rl
	r.Right = rr

	rl.Left = &Tree{
		Value:  "5",
		Parent: rl,
	}

	rl.Right = &Tree{
		Value:  "9",
		Parent: rl,
	}

	return root
}
