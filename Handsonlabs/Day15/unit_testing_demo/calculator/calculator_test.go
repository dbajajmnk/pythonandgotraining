package calculator

import (
	"errors"
	"testing"
)

func TestAdd_TableDriven(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"zero", 0, 5, 5},
		{"negative", -2, 3, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("Add(%d,%d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestDivide_ErrorOnZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrDivideByZero) {
		t.Fatalf("expected ErrDivideByZero, got %v", err)
	}
}

func TestDivide_OK(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got != 5 {
		t.Fatalf("Divide(10,2) = %d; want 5", got)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Add(123, 456)
	}
}
