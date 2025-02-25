package main

import(
	"testing"
    "github.com/rinerte/tests/functions"
)

func TestMultiply(t *testing.T) {
    result := functions.Multiply(2, 3)
    if result != 6 {
        t.Errorf("Expected 6, got %d", result)
    }	
}