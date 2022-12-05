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

		if !strings.Contains(iv, "move") {
			continue
		}

		b := strings.Split(iv, " ")
		m, _ := strconv.Atoi(b[1])
		f, _ := strconv.Atoi(b[3])
		t, _ := strconv.Atoi(b[5])
		f--
		t--

		if !part2 {
			for j := 0; j < m; j++ {
				stacks[t] = append(stacks[t], stacks[f][len(stacks[f])-1])
				stacks[f] = stacks[f][:len(stacks[f])-1]
			}
		} else {
			for j := m; j > 0; j-- {
				stacks[t] = append(stacks[t], stacks[f][len(stacks[f])-j])
			}
			stacks[f] = stacks[f][:len(stacks[f])-m]
		}
	}

	result := ""
	for k := 0; k < 9; k++ {
		result += string(stacks[k][len(stacks[k])-1])
	}

	fmt.Println(result)
}
