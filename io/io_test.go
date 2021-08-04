package io

import "testing"

func TestFunction2(t *testing.T)  {
	want := 1024 + 156
	if got := Function2(); got != want {
		t.Errorf("Function2() = %q, want %q", got, want)
	}
}

func TestReadLine(t *testing.T) {
	want := 1024 + 156
	if got := ReadLine(); got != want {
		t.Errorf("Function2() = %q, want %q", got, want)
	}
}
