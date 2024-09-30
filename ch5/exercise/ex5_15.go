package exercise

import "math"

// Exercise 5.15: Write variadic functions max and min, analogous to sum. What should these
// functions do when called with no arguments?  Write variants that require at least one argument.

// Variadic max function
func max(nums ...int) int {
	if len(nums) == 0 {
		return math.MinInt64 // Return the smallest possible integer if no arguments are provided
	}
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

// Variadic min function
func min(nums ...int) int {
	if len(nums) == 0 {
		return math.MaxInt64 // Return the largest possible integer if no arguments are provided
	}
	minVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

// Max function requiring at least one argument
func max(first int, rest ...int) int {
	maxVal := first
	for _, num := range rest {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

// Min function requiring at least one argument
func min(first int, rest ...int) int {
	minVal := first
	for _, num := range rest {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}
