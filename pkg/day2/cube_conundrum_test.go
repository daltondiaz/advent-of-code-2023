package day2

import (
	"log"
	"testing"
)

func TestSolver(t *testing.T) {
    got, _ := Solver("./input_test.txt")
    expected := 8
    
    if got != expected {
        log.Fatalf("Expected %v got %v", expected, got)
    }

    _, got= Solver("./input_test.txt")
    expected = 2286 
    
    if got != expected {
        log.Fatalf("Expected %v got %v", expected, got)
    }
}
