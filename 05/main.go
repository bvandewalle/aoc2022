package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	stacks := [9][]byte{}

	for i, iv := range input {
		if iv == "" {
			for j := i - 2; j >= 0; j-- {
				for k := 0; k < 9; k++ {
					if input[j][1+(4*k)] != ' ' {
						stacks[k] = append(stacks[k], input[j][1+(4*k)])
					}
				}
			}
		}

		if strings.Contains(iv, "move") {
			b := strings.Split(iv, " ")
			m, _ := strconv.Atoi(b[1])
			f, _ := strconv.Atoi(b[3])
			t, _ := strconv.Atoi(b[5])

			if !part2 {
				for j := 0; j < m; j++ {
					stacks[t-1] = append(stacks[t-1], stacks[f-1][len(stacks[f-1])-1])
					stacks[f-1] = stacks[f-1][:len(stacks[f-1])-1]
				}
			} else {
				for j := m; j > 0; j-- {
					stacks[t-1] = append(stacks[t-1], stacks[f-1][len(stacks[f-1])-j])
				}
				stacks[f-1] = stacks[f-1][:len(stacks[f-1])-m]
			}
		}
	}

	response := ""
	for k := 0; k < 9; k++ {
		response += string(stacks[k][len(stacks[k])-1])
	}

	fmt.Println(response)
}
