package day3

import (
	"log"
	"testing"
)

func TestSolver(t *testing.T) {
    expected := 4361
    got := Solver("./input_test.txt")
    
    if got != expected {
        log.Fatalf("Exptected %v got %v", expected, got)
    }
}
