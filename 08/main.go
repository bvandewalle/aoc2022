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

type dir struct {
	files     map[string]int
	dir       map[string]*dir
	totalSize int
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
	seen1 := 0
	for k := j - 1; k >= 0; k-- {
		seen1++
		if trees[i][k] >= trees[i][j] {
			break
		}
	}

	seen2 := 0
	for k := j + 1; k < len(trees[i]); k++ {
		seen2++
		if trees[i][k] >= trees[i][j] {
			break
		}
	}

	seen3 := 0
	for k := i - 1; k >= 0; k-- {
		seen3++
		if trees[k][j] >= trees[i][j] {
			break
		}
	}

	seen4 := 0
	for k := i + 1; k < len(trees); k++ {
		seen4++
		if trees[k][j] >= trees[i][j] {
			break
		}
	}

	return seen1 * seen2 * seen3 * seen4
}
