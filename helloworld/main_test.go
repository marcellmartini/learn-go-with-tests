package main

import (
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
	t.Run("Prime of 1", func(t *testing.T) {
		got := prime(1)
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("Prime of 3", func(t *testing.T) {
		got := prime(3)
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("Prime of 4", func(t *testing.T) {
		got := prime(4)
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("Prime of 21", func(t *testing.T) {
		got := prime(21)
		want := false
		assertCorrectMessage(t, got, want)
	})
}

func TestPrimeBetween(t *testing.T) {
	t.Run("Prime between 1 10", func(t *testing.T) {
		got := primeBetween(1, 10)
		want := []int{2, 3, 5, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
