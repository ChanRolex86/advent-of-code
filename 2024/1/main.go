package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const filename = "input.txt"

func main() {
	a, b := parseNums()
	sortNums(a)
	sortNums(b)

	// part 1: start
	diff := getTotalDifference(a, b)

	fmt.Printf("Total difference: %d\n", diff)
	// part 1: finish

	// part 2: start
	score := getSimilarityScore(a, b)

	fmt.Printf("Similarity Score: %d\n", score)
	// part 2: finish
}

func parseNums() ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var a, b []int

	var split bool
	var num int
	var space byte = ' '
	var zero byte = '0'

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		split = false
		num = 0

		lastIndex := len(line) - 1

		for index, element := range line {
			if element == space {
				if !split {
					a = append(a, num)
					num = 0
					split = true
				}
				continue
			}

			num = num*10 + int(element-zero)

			if index == lastIndex {
				b = append(b, num)
			}
		}
	}
	return a, b
}

func sortNums(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func getTotalDifference(a, b []int) int {
	total := 0

	for i, x := range a {
		total = total + getElementDifference(x, b[i])
	}

	return total
}

func getElementDifference(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func getSimilarityScore(a, b []int) int {
	score := 0

	bi := 0
	tally := 0

	for i, x := range a {
		if bi == len(b) {
			break
		}

		if i > 0 && x == a[i-1] {
			score = score + (x * tally)
			continue
		}

		tally = 0

		for bi < len(b) {
			if b[bi] > x {
				break
			}
			if b[bi] == x {
				tally = tally + 1
			}
			bi = bi + 1
		}
		score = score + (x * tally)
	}

	return score
}
