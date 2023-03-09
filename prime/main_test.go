package main

import (
	"reflect"
	"testing"
)

func TestISPrime(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{
			name:   "Prime of 1",
			number: 1,
			want:   false,
		},
		{
			name:   "Prime of 3",
			number: 3,
			want:   true,
		},
		{
			name:   "Prime of 4",
			number: 4,
			want:   false,
		},
		{
			name:   "Prime of 21",
			number: 21,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertCorrectMessage(t, tt.number, isPrime(tt.number), tt.want)
		})
	}
}

func TestPrimeBetweenNumbers(t *testing.T) {
	t.Run("Prime between 1 21", func(t *testing.T) {
		got := primeBetweenNumbers(1, 21)
		want := []int{2, 3, 5, 7, 11, 13, 17, 19}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func assertCorrectMessage(t testing.TB, number int, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("isPrime(%v) = %v, want: %v", number, got, want)
	}
}
