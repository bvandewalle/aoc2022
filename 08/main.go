package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []string

	for scanner.Scan() {
		v := scanner.Text()
		input = append(input, v)
	}

	parts(input)
}

func parts(input []string) {

	trees := [][]int{}

	for _, iv := range input {
		l := []int{}
		for _, jv := range iv {
			b, _ := strconv.Atoi(string(jv))
			l = append(l, b)
		}
		trees = append(trees, l)
	}

	count1 := 0
	count2 := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[0]); j++ {
			if isVisible(trees, i, j) {
				count1++
			}
			if a := score(trees, i, j); a > count2 {
				count2 = a
			}
		}
	}

	fmt.Println(count1)
	fmt.Println(count2)
}

func isVisible(trees [][]int, i int, j int) bool {
	seen := true

	for k := 0; k < j; k++ {
		if trees[i][k] >= trees[i][j] {
			seen = false
		}
	}
	if seen {
		return true
	}

	seen = true
	for k := j + 1; k < len(trees[i]); k++ {
		if trees[i][k] >= trees[i][j] {
			seen = false
		}
	}
	if seen {
		return true
	}

	seen = true
	for k := 0; k < i; k++ {
		if trees[k][j] >= trees[i][j] {
			seen = false
		}
	}
	if seen {
		return true
	}

	seen = true
	for k := i + 1; k < len(trees); k++ {
		if trees[k][j] >= trees[i][j] {
			seen = false
		}
	}
	return seen
}

func score(trees [][]int, i int, j int) int {
	score := 1
	seen := 0
	for k := j - 1; k >= 0; k-- {
		seen++
		if trees[i][k] >= trees[i][j] {
			break
		}
	}
	score *= seen

	seen = 0
	for k := j + 1; k < len(trees[i]); k++ {
		seen++
		if trees[i][k] >= trees[i][j] {
			break
		}
	}
	score *= seen

	seen = 0
	for k := i - 1; k >= 0; k-- {
		seen++
		if trees[k][j] >= trees[i][j] {
			break
		}
	}
	score *= seen

	seen = 0
	for k := i + 1; k < len(trees); k++ {
		seen++
		if trees[k][j] >= trees[i][j] {
			break
		}
	}
	score *= seen

	return score
}
