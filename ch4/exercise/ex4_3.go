package exercise

// Exercise 4.3: Rewrite reverse to use an array pointer in tead of a slice.
func reverse(arr *[5]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}
