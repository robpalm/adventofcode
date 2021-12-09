package aoc2021

import (
	"bufio"
	"math"
	"robpalm/adventofcode/internal/utils"
)

func Day1_1(scanner *bufio.Scanner) int {
	previous := math.MaxInt32
	increase := 0

	for scanner.Scan() {
		current := utils.ToInt(scanner.Text())
		if current > previous {
			increase++
		}
		previous = current
	}
	return increase
}

func Day1_2(scanner *bufio.Scanner) int {
	arr := make([]int, 3)
	var row, increase, current, previous int

	for scanner.Scan() {
		row++
		val := utils.ToInt(scanner.Text())
		arr = append(arr[1:], val)
		current = arr[0] + arr[1] + arr[2]

		if row > 3 && current > previous {
			increase++
		}
		previous = current
	}
	return increase
}
