package utils

import "testing"

func TestNearEqual(t *testing.T) {
	num := 3.6
	target := 3.6000001

	if !NearEqual(num, target) {
		t.Fatal("Failed")
	}
}

func TestRound(t *testing.T) {
	num := Round(3.6, 1)
	ans := 4.0

	if !NearEqual(num, ans) {
		t.Fatal("not equal")
	}
}

func TestSpecifiedRound(t *testing.T) {
	num := Round(3.666667, 6)
	ans := 3.66667

	if !NearEqual(num, ans) {
		t.Fatal("not equal")
	}
}
