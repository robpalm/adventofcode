package aoc2021

import (
	"bufio"
	"robpalm/adventofcode/internal/utils"
	"strings"
)

func DayX_1(scanner *bufio.Scanner) int {
	return DayX(scanner, true)
}

func DayX_2(scanner *bufio.Scanner) int {
	return DayX(scanner, false)
}

func DayX(scanner *bufio.Scanner, part1 bool) (result int) {
	var data []int

	for scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), ",") {
			data = append(data, utils.ToInt(s))
		}
	}
	return
}
