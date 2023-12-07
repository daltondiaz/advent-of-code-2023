package main

import (
	"daltondiaz/aoc-go-2023/pkg/day1"
	"daltondiaz/aoc-go-2023/pkg/day2"
	"daltondiaz/aoc-go-2023/pkg/day3"
	"fmt"
)

func main() {
	firstResult, secResult := day1.Solver("./pkg/day1/input.txt")
	fmt.Println("Day 1 - Part 1 - ", firstResult)
	fmt.Println("Day 1 - Part 2 - ", secResult)
    part1, part2 := day2.Solver("./pkg/day2/input.txt")
	fmt.Println("Day 2 - part 1 - ", part1)
	fmt.Println("Day 2 - part 2 - ", part2)

    resultDay3 := day3.Solver("./pkg/day3/input.txt")
	fmt.Println("Day 3 - Part 1 - ", resultDay3)
}
