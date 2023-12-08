package day3

import (
	"daltondiaz/aoc-go-2023/pkg/io"
	"fmt"
	"log"
	"strconv"
	"unicode"
)


func findFirstIdx(start int, row string) int {
	for i := start; i < len(row); i++ {
		if _, err := strconv.Atoi(string(row[i])); err == nil {
			return i
		}
	}
	return -1
}
func findLastIdx(start int, row string) int {
	for i := start; i <len(row); i++ {
		if _, err := strconv.Atoi(string(row[i])); err != nil {
			return i
		}
	}
	return len(row)
}

func evenDiagonally(first int, last int, currRow int, rows []string, adj [][]int) bool {

    // walk adj
    for i:= first; i<last; i++ {
        for j:=0; j<len(adj); j++ {
            y := currRow + adj[j][0]
            x := i + adj[j][1] 

            // base cases
            // 1. verify if is not out bound of rows
            if y < 0 || x < 0 || y == len(rows) || x == len(rows[y]) {
                continue
            }

            // 2. not check if number in the middle are adj 
            // Ex:      ..1234*..
            //       check: 2
            // jump on 1 and 3
            if y == currRow && (x >= first && x<last) {
                continue
            }

            if string(rows[y][x]) != "." && (unicode.IsSymbol(rune(rows[y][x])) || unicode.IsPunct(rune(rows[y][x]))) {
                // fmt.Println(rows[currRow][first:last]) 
                return true 
            }
        }
    }
    return false
}

func Solver(path string) int {
	rows, err := io.Read(path)
    sum := 0
    // check all direction by sum y,x
    adj := [][]int{{-1,0},{-1,1},{0,1},{1,1},{1,0},{1,-1},{0,-1},{-1,-1}}
	if err != nil {
		log.Fatalf("Error to read input file")
	}
	for i := 0; i < len(rows); i++ {
		j := 0
        first := 0
        last := 0
		for {
			if j >= len(rows[i]) {
				break
			}
			first = findFirstIdx(j, rows[i])
            if first < 0 {
                j++
                continue
            }
			last = findLastIdx(first, rows[i])
            if last < 0 {
                j++
                continue
            }
			if first > 0 || last > 0 {
                if evenDiagonally(first, last, i, rows, adj) {
                    v, err := strconv.Atoi(rows[i][first:last])
                    if err != nil {
                        fmt.Println("Error to convert")
                    }
                    sum += v
                }
			}
			j = last + 1
		}

	}
	return sum 
}
