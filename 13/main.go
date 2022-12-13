package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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
	pairs := [][]any{}
	currentPair := []any{}

	for _, iv := range input {
		if iv == "" {
			pairs = append(pairs, currentPair)
			currentPair = []any{}
			continue
		}

		var l any
		json.Unmarshal([]byte(iv), &l)
		currentPair = append(currentPair, l)
	}
	pairs = append(pairs, currentPair)

	if !part2 {
		ans := 0
		for i, p := range pairs {
			if recurCompare(p[0], p[1]) > 0 {
				ans += i + 1
			}
		}
		fmt.Println(ans)
		return
	}

	p2r, p6r := 0, 0
	for _, p := range pairs {
		for _, pp := range p {
			if recurCompare(pp, []any{[]any{float64(2)}}) > 0 {
				p2r++
			}
			if recurCompare(pp, []any{[]any{float64(6)}}) > 0 {
				p6r++
			}
		}
	}
	fmt.Println((p2r + 1) * (p6r + 2))
}

func recurCompare(l, r any) int {
	tl := reflect.TypeOf(l)
	tr := reflect.TypeOf(r)

	// Both List
	if tl.Kind() == reflect.Slice && tr.Kind() == reflect.Slice {
		la := l.([]any)
		ra := r.([]any)
		for i, il := range la {
			if i < len(ra) {
				r := recurCompare(il, ra[i])
				if r != 0 {
					return r
				}
			} else {
				return -1
			}
		}
		if len(la) < len(ra) {
			return 1
		} else if len(la) > len(ra) {
			return -1
		}
		return 0
	}

	// One List
	if tl.Kind() == reflect.Slice {
		return recurCompare(l, []any{r})
	}
	if tr.Kind() == reflect.Slice {
		return recurCompare([]any{l}, r)
	}

	// Both Integers
	la := l.(float64)
	ra := r.(float64)
	if la < ra {
		return 1
	} else if la > ra {
		return -1
	}
	return 0
}
