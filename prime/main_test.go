package main

import (
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
	t.Run("Prime of 1", func(t *testing.T) {
		got := isPrime(1)
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("Prime of 3", func(t *testing.T) {
		got := isPrime(3)
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("Prime of 4", func(t *testing.T) {
		got := isPrime(4)
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("Prime of 21", func(t *testing.T) {
		got := isPrime(21)
		want := false
		assertCorrectMessage(t, got, want)
	})
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

func assertCorrectMessage(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
