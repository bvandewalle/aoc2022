package main

import (
	"fmt"
	"sort"
)

func main() {
	parts(false)
	parts(true)
}

func parts(part2 bool) {
	items := [][]int{{62, 92, 50, 63, 62, 93, 73, 50}, {51, 97, 74, 84, 99}, {98, 86, 62, 76, 51, 81, 95}, {53, 95, 50, 85, 83, 72}, {59, 60, 63, 71}, {92, 65}, {78}, {84, 93, 54}}
	ops := []func(int) int{createOp(7, false, false), createOp(3, true, false), createOp(4, true, false), createOp(5, true, false), createOp(5, false, false), createOp(0, false, true), createOp(8, true, false), createOp(1, true, false)}
	tests := []int{2, 7, 13, 19, 11, 5, 3, 17}
	results := [][]int{{7, 1}, {2, 4}, {5, 4}, {6, 0}, {5, 3}, {6, 3}, {0, 7}, {2, 1}}

	inspections := []int{}
	for i := 0; i < 8; i++ {
		inspections = append(inspections, 0)
	}

	CommonMult := 1
	for _, iv := range tests {
		CommonMult *= iv
	}

	rounds := 10
	if part2 {
		rounds = 10000
	}

	for r := 0; r < rounds; r++ {
		for i, it := range items {
			for _, jv := range it {
				new := ops[i](jv)
				if !part2 {
					new = new / 3
				} else {
					new %= CommonMult
				}
				if new%tests[i] == 0 {
					items[results[i][0]] = append(items[results[i][0]], new)
				} else {
					items[results[i][1]] = append(items[results[i][1]], new)
				}
				inspections[i]++
			}
			items[i] = []int{}
		}
	}

	sort.Ints(inspections)
	fmt.Println(inspections[6] * inspections[7])
}

func createOp(x int, opAdd bool, same bool) func(y int) int {
	if opAdd {
		return func(y int) int {
			return x + y
		}
	}

	if same {
		return func(y int) int {
			return y * y
		}
	}

	return func(y int) int {
		return x * y
	}
}
