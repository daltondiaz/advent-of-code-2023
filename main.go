package main

import (
	"daltondiaz/aoc-go-2023/pkg/day1"
	"daltondiaz/aoc-go-2023/pkg/day2"
	"fmt"
)

func main() {
	firstResult, secResult := day1.Solver("./pkg/day1/input.txt")
	fmt.Println("Day 1 - Part 1 - ", firstResult)
	fmt.Println("Day 1 - Part 2 - ", secResult)
	fmt.Println("Day 2 - ", day2.Solver("./pkg/day2/input.txt"))
}
