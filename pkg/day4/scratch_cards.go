package day4

import (
	"daltondiaz/aoc-go-2023/pkg/io"
	"log"
	"strings"
)


func double(num int) int {
    if num == 0 {
        return 0 
    }
    if num == 1 {
        return num
    }
    result := 1
    for i:=num; i>= 2; i-- {
        result *= 2
    }
    return result
}

func hasNum(nums []string, target string) bool {
    for _, v :=  range nums {
        if strings.TrimSpace(v) == "" {
            continue
        }
        if strings.TrimSpace(v) == strings.TrimSpace(target)  {
            return true
        }
    }
    return false
}

func walk(rows []string) int {
    sum := 0
    for i:= 0; i<len(rows); i++ {
        line := strings.Trim(rows[i], " ")
        card := strings.Split(line, ":")
        values := strings.Split(card[1], "|")
        winNum := strings.Split(values[0], " ")
        myNum := strings.Split(values[1], " ")
        count := 0
        for _, v := range winNum {
            if hasNum(myNum, v) {
                count ++
            }
        }
        result := double(count)
        sum += result 
    }
    return sum
}

func Part1(path string) int {
	rows, err := io.Read(path)
	if err != nil {
		log.Fatalf("Error to convert input file")
	}
    return walk(rows)
}
