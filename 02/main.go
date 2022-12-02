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

	for _, iv := range input {
		a := 2 - int('C'-iv[0])
		b := 2 - int('Z'-iv[2])

		count += 3 * ((b - a + 4) % 3)
		count += b + 1
	}

	fmt.Println(count)
}

func part2(input []string) {
	count := 0

	for _, iv := range input {
		a := 2 - int('C'-iv[0])
		b := 2 - int('Z'-iv[2])

		count += b * 3
		count += 1 + ((a + b + +2) % 3)
	}

	fmt.Println(count)
}
