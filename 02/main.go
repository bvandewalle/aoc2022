package main

import (
	"bufio"
	"fmt"
	"os"
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

	part1(input)
	part2(input)
}

func part1(input []string) {
	count := 0
	points := [][]int{[]int{3, 6, 0}, []int{0, 3, 6}, []int{6, 0, 3}}

	for _, iv := range input {
		a := 2 - int('C'-iv[0])
		b := 2 - int('Z'-iv[2])

		count += points[a][b]
		count += b + 1
	}

	fmt.Println(count)
}

func part2(input []string) {
	count := 0
	points := [][]int{[]int{3, 1, 2}, []int{1, 2, 3}, []int{2, 3, 1}}

	for _, iv := range input {
		a := 2 - int('C'-iv[0])
		b := 2 - int('Z'-iv[2])

		count += b * 3
		count += points[a][b]
		fmt.Println(count)

	}

	fmt.Println(count)
}
