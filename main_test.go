package main

import "testing"

func TestAbs(t *testing.T) {
	got := 2
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
