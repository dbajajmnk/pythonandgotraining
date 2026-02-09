// File: mathops_test.go
package mathops

import (
	"fmt"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	fmt.Println("N ", b.N)
	for i := 0; i < b.N; i++ {

		_ = Add(10, 20)
	}
}

func BenchmarkAddWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AddWithLoop(10, 20)
	}
}

func BenchmarkAdd_Sub(b *testing.B) {
	cases := []struct {
		name string
		a, c int
	}{
		{"small", 1, 2},
		{"medium", 100, 200},
		{"large", 1000, 2000},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Add(tc.a, tc.c)
			}
		})
	}
}
