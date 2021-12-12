package aoc2021

import (
	"bufio"
	"math"
	"robpalm/adventofcode/internal/utils"
	"strings"
)

func Day7_1(scanner *bufio.Scanner) int {
	return Day7(scanner, true)
}

func Day7_2(scanner *bufio.Scanner) int {
	return Day7(scanner, false)
}

func Day7(scanner *bufio.Scanner, part1 bool) int {
	var crabPos []int

	for scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), ",") {
			crabPos = append(crabPos, utils.ToInt(s))
		}
	}

	var maxPos int
	for _, v := range crabPos {
		if v > maxPos {
			maxPos = v
		}
	}

	minFuel := math.MaxInt
	for i := 0; i < maxPos; i++ {
		fuel := 0
		for _, v := range crabPos {
			steps := utils.Abs(i - v)
			if part1 {
				fuel += steps
			} else {
				if steps%2 == 0 {
					fuel += (steps + 1) * steps / 2
				} else {
					fuel += (steps)*(steps-1)/2 + steps
				}
			}
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}
