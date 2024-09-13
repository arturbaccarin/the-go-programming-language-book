package exercise

/*
Exercise 3.12: Write a function that reports whether two strings are anagrams of each other,
that is, they contain the same letters in a different order.
*/

func isAnagram(s1, s2 string) bool {
	countS1 := make(map[rune]int)
	countS2 := make(map[rune]int)

	if len(s1) != len(s2) {
		return false
	}

	for _, c := range s1 {
		countS1[c]++
	}

	for _, c := range s2 {
		countS2[c]++
	}

	for k, v := range countS1 {
		if countS2[k] != v {
			return false
		}
	}

	return true
}
