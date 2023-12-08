package day3

import (
	"log"
	"testing"
)

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

func TestLucas(t *testing.T) {

    expected := 6719
    got := Solver("./input_lucas.txt")
    
    if got != expected {
        log.Fatalf("Exptected %v got %v", expected, got)
    }
}
