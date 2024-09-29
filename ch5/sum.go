package ch5

/*
To declare a var iadic function, the typ e of the ﬁnal parameter is pre ceded by an ellipsis, ‘‘...’’,
which indic ates that the function may be cal le d with any number of arguments of this typ e.
*/
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
