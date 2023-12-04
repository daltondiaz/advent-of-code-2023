package day2

import (
	"log"
	"testing"
)

func TestSolver(t *testing.T) {
    got := Solver("./input_test.txt")
    expected := 8
    
    if got != expected {
        log.Fatalf("Expected %v got %v", expected, got)
    }
}
