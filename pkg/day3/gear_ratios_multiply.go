package day3

import (
	"daltondiaz/aoc-go-2023/pkg/io"
	"log"
	"strconv"
	"unicode"
)

var visited [][]bool

func walkBack(next int, currRow int, rows []string) int{
    startIdx := next 
    for  {
        if next < 0 {
            startIdx = 0 
            next = 0
            break
        }
        if visited[currRow][next] {
            next -- 
            continue
        }
        if !unicode.IsDigit(rune(rows[currRow][next]))  && !visited[currRow][next] {
            break
        }
        startIdx=next
        next--
    }
    return startIdx
}

func walkForward(next int, currRow int, rows []string) int {
    endIdx := next
    for {
        if next == len(rows[currRow]) {
            break
        }
        if visited[currRow][next] {
            next ++
            continue
        }
        if !unicode.IsDigit(rune(rows[currRow][next]))  && !visited[currRow][next] {
            break
        }
        endIdx = next
        next++
    }
    return endIdx
}

func walk(rows []string, x,y int) int {
    next := x
    startIdx := walkBack(next, y, rows)
    endIdx := walkForward(next, y, rows)
    result, err := strconv.Atoi(rows[y][startIdx:endIdx+1]) 
    if err != nil {
        log.Fatal("error to convert", err)
    }
    setVisited(y, startIdx, endIdx)
    return result 
}

func setVisited(y, start, end int) {
    for i:=start; i <= end; i++ {
        visited[y][i] = true
    }
}

func multiplyAdjacent(idxSign int, currRow int, rows []string, adj [][]int) int {

	// walk adj
    var numbers []int
	for j := 0; j < len(adj); j++ {
		y := currRow + adj[j][0]
		x := idxSign + adj[j][1]

		// base cases
		// 1. verify if is not out bound of rows
		if y < 0 || x < 0 || y == len(rows) || x == len(rows[y]) {
			continue
		}

		if unicode.IsDigit(rune(rows[y][x])) && !visited[y][x] {
            numbers = append(numbers, walk(rows,x,y))
		}
	}

    if len(numbers) == 1 {
        return 0
    }
    result := 1 
    for i :=0; i<len(numbers); i++ {
        result *= numbers[i]
    }
    
	return result
}

func Multiply(path string) int {
	rows, err := io.Read(path)
	if err != nil {
		log.Fatalf("Error to read input file")
	}
	adj := [][]int{{-1,0},{-1,1},{0,1},{1,1},{1,0},{1,-1},{0,-1},{-1,-1}}

    visited = make([][]bool, len(rows))
    for i := range visited {
        visited[i] = make([]bool, len(rows[0]))
    }

    result := 0
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[y]); x++ {
			if string(rows[y][x]) == "*" {
                result += multiplyAdjacent(x, y, rows, adj)
			}
		}
	}
	return result 
}
