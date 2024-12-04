package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "input.txt"
const space byte = ' '
const zero byte = '0'

func main() {
	reports := parseReports()

	// part 1: start
	safeCount := getSafeCount(reports)

	fmt.Printf("No. of safe reports: %d\n", safeCount)
	// part 1: finish

	// part 2: start
	dampenedSafeCount := getDampenedSafeCount(reports)

	fmt.Printf("No. of dampened safe reports: %d\n", dampenedSafeCount)
	// part 2: finish
}

func parseReports() [][]int {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	var x [][]int
	var temp []int
	var num int

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		num = 0
		temp = nil

		lastIndex := len(line) - 1

		for index, element := range line {
			if element == space {
				temp = append(temp, num)
				num = 0
				continue
			}

			num = num*10 + int(element-zero)

			if index == lastIndex {
				temp = append(temp, num)
				x = append(x, temp)
			}
		}
	}

	return x
}

func getSafeCount(reports [][]int) int {
	safe := 0

	for _, x := range reports {
		if isReportSafe(x) {
			safe = safe + 1
		}
	}

	return safe
}

func isReportSafe(report []int) bool {
	if report[0] == report[1] {
		return false
	}

	var mp int
	if report[1] > report[0] {
		mp = 1
	} else {
		mp = -1
	}

	for i, x := range report[1:] {
		if (mp*(x-report[i]) < 1) || (mp*(x-report[i]) > 3) {
			return false
		}
	}

	return true
}

func getDampenedSafeCount(reports [][]int) int {
	safe := 0

	for _, x := range reports {
		if isReportDampenedSafe(x) {
			safe = safe + 1
		}
	}

	return safe
}

func isReportDampenedSafe(report []int) bool {
	failingIndex := getFailinigIndex(report)
	if failingIndex == -1 {
		return true
	} else {
		var i = 0
		for i < failingIndex+2 {
			if isReportSafe(excludeIndex(report, i)) {
				return true
			}
			i = i + 1
		}
		return false
	}
}

func getFailinigIndex(report []int) int {
	if report[0] == report[1] {
		return 1
	}

	var mp int
	if report[1] > report[0] {
		mp = 1
	} else {
		mp = -1
	}

	for i, x := range report[1:] {
		if (mp*(x-report[i]) < 1) || (mp*(x-report[i]) > 3) {
			return i + 1
		}
	}

	return -1
}

func excludeIndex(slice []int, index int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:index])
	copy(newSlice[index:], slice[index+1:])
	return newSlice
}
