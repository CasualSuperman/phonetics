package sift3

import "testing"

func TestEqualSift(t *testing.T) {
	s1 := "Hello, world."
	s2 := "Hello, world."

	if Sift(s1, s2) != 0 {
		t.Error("Distance between equal strings should be zero.")
	}
}

func TestTypoSift(t *testing.T) {
	s1 := "Hello, world."
	s2 := "Hello! world."

	if Sift(s1, s2) != 1 {
		t.Error("Distance between two strings with one typo should be one.")
	}
}

func TestDropSift(t *testing.T) {
	s1 := "Hello, world."
	s2 := "Hello world."

	if Sift(s1, s2) != 1.5 {
		t.Error("Distance between two strings with a missing character should be 1.5")
	}
}
