package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("main")

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var input []string

	for scanner.Scan() {
		v := scanner.Text()
		input = append(input, v)
	}

	file.Close()

	part1(input)
	part2(input)
}

func part1(input []string) {
	maxCount := 0

	count := 0
	for _, iv := range input {
		if iv == "" {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
		v, _ := strconv.Atoi(iv)
		count += v
	}

	fmt.Println(maxCount)
}

func part2(input []string) {
	allCounts := []int{}

	count := 0
	for _, iv := range input {
		if iv == "" {
			allCounts = append(allCounts, count)
			count = 0
		}
		v, _ := strconv.Atoi(iv)
		count += v
	}

	sort.Slice(allCounts, func(i, j int) bool {
		return allCounts[i] > allCounts[j]
	})

	fmt.Println(allCounts[0] + allCounts[1] + allCounts[2])
}
