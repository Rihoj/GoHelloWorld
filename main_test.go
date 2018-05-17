package main

import (
	"testing"
)

// TestPrint makes sure that the print function works as expected
func TestPrint(t *testing.T) {
	var output = print()
	var expected = "Hello World!"
	if output != expected {
		t.Errorf("Out was incorrect, got: %s, want: %s", output, expected)
	}
}
func TestMain(t *testing.T) {
	main()
}
func BenchmarkTestPrint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		print()
	}
}
