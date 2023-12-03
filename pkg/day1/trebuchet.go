package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func replaceDigit(text string) string {
	curr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	next := []string{"on1e", "tw2o", "thr3ee", "fo4ur", "fi5ve", "si6x", "sev7en", "ei8ght", "ni9ne"}

	for i := 0; i < len(next); i++ {
		text = strings.ReplaceAll(text, curr[i], next[i])
	}

	return text
}

func findFirstAndLastDigit(text string, isReplace bool) int {
	var result strings.Builder
	var newText string
	if isReplace {
		newText = replaceDigit(text)
	} else {
		newText = text
	}
	i := 0
	for {
		if i == len(newText) {
			break
		}

		if _, err := strconv.Atoi(string(newText[i])); err == nil {
			result.WriteString(string(newText[i]))
			break
		}
		i++
	}

	i = len(newText) - 1
	for {
		if i < 0 {
			break
		}
		if _, err := strconv.Atoi(string(newText[i])); err == nil {
			result.WriteString(string(newText[i]))
			break
		}
		i--
	}
	number, err := strconv.Atoi(result.String())
	if err != nil {
		return 0
	}
	return number
}

func Solver(path string) (int, int) {
	lines, err := readInput(path)
	if err != nil {
		log.Fatal("error to read the file")
	}

	var part1 int
	var part2 int
	for i := 0; i < len(lines); i++ {
		r := int(findFirstAndLastDigit(lines[i], false))
		part1 += r
		replaced := int(findFirstAndLastDigit(lines[i], true))
        part2 += replaced
	}
	return part1, part2 
}
