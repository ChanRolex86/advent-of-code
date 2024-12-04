package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "input.txt"
const req = "mul(,)"
const zero = '0'
const do = "do()"
const dont = "don't()"

func main() {
	// part 1: start
	x1 := processFilePart1()
	sum1 := sumArray(x1)

	fmt.Printf("Total sum: %d\n", sum1)
	// part 1: finish

	// part 2: start
	x2 := processFilePart2()
	sum2 := sumArray(x2)

	fmt.Printf("Total sum: %d\n", sum2)
	// part 2: finish
}

func processFilePart1() []int {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	var x []int
	var reqIndex, num1, num2 = 0, 0, 0

	for {
		reqIndex, num1, num2 = 0, 0, 0

		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		for _, element := range line {
			if element == req[reqIndex] {
				if element == ')' {
					if num1 != 0 && num2 != 0 {
						x = append(x, num1*num2)
					}
					reqIndex, num1, num2 = 0, 0, 0
				} else {
					reqIndex = reqIndex + 1
				}
			} else if req[reqIndex] == ',' {
				if zero <= element && element <= '9' {
					num1 = num1*10 + int(element-zero)
				} else {
					reqIndex, num1, num2 = 0, 0, 0
				}
			} else if req[reqIndex] == ')' {
				if zero <= element && element <= '9' {
					num2 = num2*10 + int(element-zero)
				} else {
					reqIndex, num1, num2 = 0, 0, 0
				}
			} else {
				reqIndex, num1, num2 = 0, 0, 0
			}
		}
	}

	return x
}

func processFilePart2() []int {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	var x []int
	var reqIndex, num1, num2, commandIndex = 0, 0, 0, 0
	var isDo = true

	for {
		reqIndex, num1, num2 = 0, 0, 0

		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		for _, element := range line {
			if isDo {
				if element == req[reqIndex] {
					if element == ')' {
						if num1 != 0 && num2 != 0 {
							x = append(x, num1*num2)
						}
						reqIndex, num1, num2 = 0, 0, 0
					} else {
						reqIndex = reqIndex + 1
					}
				} else if req[reqIndex] == ',' {
					if zero <= element && element <= '9' {
						num1 = num1*10 + int(element-zero)
					} else {
						reqIndex, num1, num2 = 0, 0, 0
					}
				} else if req[reqIndex] == ')' {
					if zero <= element && element <= '9' {
						num2 = num2*10 + int(element-zero)
					} else {
						reqIndex, num1, num2 = 0, 0, 0
					}
				} else {
					reqIndex, num1, num2 = 0, 0, 0
				}
				if reqIndex == 0 && element == dont[commandIndex] {
					if element == ')' {
						isDo = false
						commandIndex = 0
					} else {
						commandIndex = commandIndex + 1
					}
				} else {
					commandIndex = 0
				}
			} else {
				if element == do[commandIndex] {
					if element == ')' {
						isDo = true
						commandIndex = 0
					} else {
						commandIndex = commandIndex + 1
					}
				} else {
					commandIndex = 0
				}
			}
		}
	}

	return x
}

func sumArray(x []int) int {
	sum := 0

	for _, element := range x {
		sum = sum + element
	}

	return sum
}
