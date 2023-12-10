package day3

import (
	"log"
	"testing"
)
func TestMutiply(t *testing.T) {

    expected := 467835 
    got := Multiply("./input_test_multiply.txt")
    
    if got != expected {
        log.Fatalf("Exptected %v got %v", expected, got)
    }
}

func TestSolverNumberEndOfLine(t *testing.T) {
    expected := 1374
    got := Solver("./input_test_2.txt")
    
    if got != expected {
        log.Fatalf("Exptected %v got %v", expected, got)
    }
}

func TestSolver(t *testing.T) {
    expected := 4361
    got := Solver("./input_test.txt")
    
    if got != expected {
        log.Fatalf("Exptected %v got %v", expected, got)
    }
}
