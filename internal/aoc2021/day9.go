package aoc2021

import (
	"bufio"
	"robpalm/adventofcode/internal/utils"
)

func Day9_1(scanner *bufio.Scanner) int {
	return Day9(scanner, true)
}

func Day9_2(scanner *bufio.Scanner) int {
	return Day9(scanner, false)
}

var rows []string

func Day9(scanner *bufio.Scanner, part1 bool) (result int) {

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	for y := range rows {
		for x := range rows[y] {
			curHeight := getHeight(y, x)

			lowHeight := utils.Min(
				getHeight(y-1, x-1), getHeight(y-1, x+0), getHeight(y-1, x+1),
				getHeight(y+0, x-1), getHeight(y+0, x+1),
				getHeight(y+1, x-1), getHeight(y+1, x+0), getHeight(y+1, x+1),
			)

			if curHeight < lowHeight {
				result += curHeight + 1
			}
		}
	}
	return
}

func getHeight(y int, x int) int {
	if y >= 0 && x >= 0 && y < len(rows) && x < len(rows[y]) {
		return utils.ToInt(string(rows[y][x]))
	}
	return 9
}
