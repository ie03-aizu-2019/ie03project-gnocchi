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

func TestRound2(t *testing.T) {
	num := Round(6.108815, 6)
	ans := 6.10882

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

func TestCeil(t *testing.T) {
	num := Ceil(6.108813, 6)
	ans := 6.10882

	if !NearEqual(num, ans) {
		t.Fatal(num, ans)
	}
}
