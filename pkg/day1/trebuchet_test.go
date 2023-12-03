package day1

import (
	"log"
	"testing"
)

func TestSolver(t *testing.T){
    got, _ := Solver("./input_test.txt")
    expected := 142
    if got != expected {
        log.Fatalf("Expected %v got %v", expected, got)
    }
}

func TestPart2(t *testing.T){
    _, got := Solver("./input_test_p2.txt")
    expected := 281
    if got != expected {
        log.Fatalf("Expected %v got %v", expected, got)
    }

}
