package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSub(t *testing.T) {
	total := sub(5, 5)
	if total != 0 {
		t.Errorf("Sub was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestMult(t *testing.T) {
	total := mult(5, 5)
	if total != 25 {
		t.Errorf("Mult was incorrect, got: %d, want: %d.", total, 25)
	}
}

func TestDiv(t *testing.T) {
	total := div(5, 5)
	if total != 1 {
		t.Errorf("Div was incorrect, got: %d, want: %d.", total, 1)
	}
}

func TestMod(t *testing.T) {
	total := mod(5, 5)
	if total != 0 {
		t.Errorf("Mod was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestPow(t *testing.T) {
	total := pow(5, 5)
	if total != 0 {
		t.Errorf("Pow was incorrect, got: %d, want: %d.", total, 0)
	}
}
