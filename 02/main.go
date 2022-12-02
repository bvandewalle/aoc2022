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
		switch a := iv[0]; a {
		case 'A':
			switch b := iv[2]; b {
			case 'X':
				count += 3 + 1
			case 'Y':
				count += 6 + 2
			case 'Z':
				count += 0 + 3
			}
		case 'B':
			switch b := iv[2]; b {
			case 'X':
				count += 0 + 1
			case 'Y':
				count += 3 + 2
			case 'Z':
				count += 6 + 3
			}
		case 'C':
			switch b := iv[2]; b {
			case 'X':
				count += 6 + 1
			case 'Y':
				count += 0 + 2
			case 'Z':
				count += 3 + 3
			}
		}
	}

	fmt.Println(count)
}

func part2(input []string) {
	count := 0

	for _, iv := range input {
		switch a := iv[0]; a {
		case 'A':
			switch b := iv[2]; b {
			case 'X':
				count += 0 + 3
			case 'Y':
				count += 3 + 1
			case 'Z':
				count += 6 + 2
			}
		case 'B':
			switch b := iv[2]; b {
			case 'X':
				count += 0 + 1
			case 'Y':
				count += 3 + 2
			case 'Z':
				count += 6 + 3
			}
		case 'C':
			switch b := iv[2]; b {
			case 'X':
				count += 0 + 2
			case 'Y':
				count += 3 + 3
			case 'Z':
				count += 6 + 1
			}
		}
		fmt.Println(count)
	}

	fmt.Println(count)
}
