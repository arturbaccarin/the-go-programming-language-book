package exercise

import (
	"testing"

	"github.com/arturbaccarin/the-go-programming-language-book/ch1"
)

/*
Exercise 1.3: Experiment to measure the difference in running time between our potentially
inefficient versions and the one that uses strings.Join.
*/
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch1.Echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch1.Echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch1.Echo3()
	}
}
