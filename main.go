package main

import (
	"daltondiaz/aoc-go-2023/pkg/day1"
	"fmt"
)

func main() {
	firstResult, secResult := day1.Solver("./pkg/day1/input.txt")
	fmt.Println("Day 1 - Part 1 - ", firstResult)
	fmt.Println("Day 1 - Part 2 - ", secResult)
}
