package exercise

// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
// RemoveAdjacentDuplicates removes adjacent duplicate strings from a slice.

// []string{"z", "a", "b", "b", "c", "a", "a", "a", "b"}
func RemoveAdjacentDuplicates(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	// Initialize a new length for the unique elements
	uniqueLen := 1

	// Iterate through the slice starting from the second element
	for i := 1; i < len(slice); i++ {
		// If the current element is different from the last unique element
		if slice[i] != slice[uniqueLen-1] {
			slice[uniqueLen] = slice[i]
			uniqueLen++
		}
	}

	// Truncate the slice to the length of unique elements
	return slice[:uniqueLen]
}
