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

	partsOptimized(input, false)
	partsOptimized(input, true)
}

func partsBrute(input []string, part2 bool) {
	limit := 4
	if part2 {
		limit = 14
	}

Exit:
	for i, _ := range input[0] {
		if i > limit-1 {
			for j := 0; j < limit; j++ {
				for k := 0; k < limit; k++ {
					if j != k {
						if input[0][i-j] == input[0][i-k] {
							continue Exit
						}
					}
				}
			}
			fmt.Println(i + 1)
			return
		}
	}
}

func partsOptimized(input []string, part2 bool) {
	le := 4
	if part2 {
		le = 14
	}

	m := map[byte]int{}

	for i, iv := range input[0] {
		m[byte(iv)] += 1

		if i > le-1 {
			m[input[0][i-le]] -= 1
			if v := m[input[0][i-le]]; v == 0 {
				delete(m, input[0][i-le])
			}
		}

		if len(m) == le {
			fmt.Println(i + 1)
			return
		}
	}
}
