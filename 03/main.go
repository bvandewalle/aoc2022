package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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
		m := map[rune]bool{}
		counted := map[rune]bool{}
		for i, c := range iv {
			if i < len(iv)/2 {
				m[c] = true
			} else {
				if _, exists := m[c]; exists {
					if _, exists2 := counted[c]; !exists2 {
						if unicode.IsLower(c) {
							count += 26 - ('z' - int(c))
						} else {
							count += 52 - ('Z' - int(c))
						}
						counted[c] = true
					}
				}
			}
		}
	}

	fmt.Println(count)
}

func part2(input []string) {
	count := 0
	m := []map[rune]bool{}

	for i, iv := range input {
		if i%3 == 0 {
			m = []map[rune]bool{}
		}
		currentLine := make(map[rune]bool)
		for _, c := range iv {
			currentLine[c] = true
		}
		m = append(m, currentLine)
		if i%3 == 2 {
			for k := range m[0] {
				if _, e1 := m[1][k]; e1 {
					if _, e2 := m[2][k]; e2 {
						if unicode.IsLower(k) {
							count += 26 - ('z' - int(k))
						} else {
							count += 52 - ('Z' - int(k))
						}
					}
				}
			}
		}
	}

	fmt.Println(count)
}
