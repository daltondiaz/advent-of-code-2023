package day2

import (
	"daltondiaz/aoc-go-2023/pkg/io"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var bagCubes = map[string]int{"red": 12, "green": 13, "blue": 14}

// Return two string with digits and letters of string text
func parser(text string) (string, string) {
	var digits, letter []rune
	for _, r := range text {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
		if unicode.IsLetter(r) {
			letter = append(letter, r)
		}
	}
	return string(digits), string(letter)
}

func part1(data []string) int {

	sumGames := 0
	for _, row := range data {
		game := strings.Split(row, ":")
		moves := strings.Split(game[1], ";")
		isGameInvalid := false

		for i := 0; i < len(moves) && !isGameInvalid; i++ {
			items := strings.Split(moves[i], ",")
			for j := 0; j < len(items) && !isGameInvalid; j++ {
				num, cube := parser(items[j])
				if v, _ := strconv.Atoi(num); v > bagCubes[cube] {
					isGameInvalid = true
				}
			}
		}
		if !isGameInvalid {
			n, _ := parser(game[0])
			v, err := strconv.Atoi(n)
			if err != nil {
				log.Printf("error to convert game %v", game[0])
			}
			sumGames += v
		}
	}

	return sumGames
}

func part2(data []string) int {

	sumGames := 0
	for _, row := range data {
		game := strings.Split(row, ":")
		moves := strings.Split(game[1], ";")
		currRed := 0
		currGreen := 0
		currBlue := 0

		for i := 0; i < len(moves); i++ {
			items := strings.Split(moves[i], ",")
			for j := 0; j < len(items); j++ {
				num, cube := parser(items[j])
				v, _ := strconv.Atoi(num)
				if cube == "green" && currGreen < v {
					currGreen = v
				}
				if cube == "red" && currRed < v {
					currRed = v
				}
				if cube == "blue" && currBlue < v {
					currBlue = v
				}
			}
		}
		sumGames += (currRed * currGreen * currBlue)
	}
	return sumGames
}

func Solver(path string) (int, int) {
	data, err := io.Read(path)

	if err != nil {
		log.Fatalf("Error to read input file")
	}
	part1 := part1(data)
	part2 := part2(data)
	return part1, part2
}
