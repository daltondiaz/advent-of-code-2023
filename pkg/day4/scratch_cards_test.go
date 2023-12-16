package day4

import (
	"log"
	"testing"
)

func TestPart1(t *testing.T){
    expected := 13
    got := Part1("./input_test.txt")

    if got != expected {
        log.Fatalf("Exptected %v got %v", expected, got)
    }
}
