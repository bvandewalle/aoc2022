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

	parts(input, false)
	parts(input, true)
}

func parts(input []string, part2 bool) {
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
