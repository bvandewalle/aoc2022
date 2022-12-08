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

	parts(input)
}

type dir struct {
	files     map[string]int
	dir       map[string]*dir
	totalSize int
}

func parts(input []string) {
	root := &dir{
		files: map[string]int{},
		dir:   map[string]*dir{},
	}

	current := root
	allDirs := []*dir{root}

	for _, iv := range input {
		if strings.Contains(iv, "/") || iv == "$ ls" {
			continue
		}

		if strings.HasPrefix(iv, "$ cd") {
			a := strings.Split(iv, " ")
			current = current.dir[a[2]]
			continue
		}

		if strings.HasPrefix(iv, "dir ") {
			a := strings.Split(iv, " ")
			newDir := &dir{
				files: map[string]int{},
				dir:   map[string]*dir{},
			}
			newDir.dir[".."] = current
			current.dir[a[1]] = newDir
			allDirs = append(allDirs, newDir)
			continue
		}
		a := strings.Split(iv, " ")
		b, _ := strconv.Atoi(a[0])
		current.files[a[1]] = b
	}

	calcSize(root)

	count := 0
	for _, v := range allDirs {
		if v.totalSize <= 100000 {
			count += v.totalSize
		}
	}
	fmt.Println(count)

	neededSpace := root.totalSize - 40000000
	toDelete := 70000000

	for _, v := range allDirs {
		if v.totalSize >= neededSpace && v.totalSize < toDelete {
			toDelete = v.totalSize
		}
	}

	fmt.Println(toDelete)

}

func calcSize(d *dir) int {
	for _, v := range d.files {
		d.totalSize += v
	}
	for k, v := range d.dir {
		if k != ".." {
			d.totalSize += calcSize(v)
		}
	}
	return d.totalSize
}
