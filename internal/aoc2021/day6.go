package aoc2021

import (
	"bufio"
	"robpalm/adventofcode/internal/utils"
	"strings"
)

type Fish struct {
	timer int
}

func Day6_1(scanner *bufio.Scanner) int {
	return Day6(scanner, 80)
}

func Day6_2(scanner *bufio.Scanner) int {
	return Day6(scanner, 256)
}

func Day6(scanner *bufio.Scanner, days int) int {
	currentFishes := make(map[int]int)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		for _, s := range parts {
			currentFishes[utils.ToInt(s)]++
		}
	}

	for i := 0; i < days; i++ {
		nextGenFishes := make(map[int]int)
		for timer, fishes := range currentFishes {
			if timer == 0 {
				nextGenFishes[6] += fishes
				nextGenFishes[8] += fishes
			} else {
				nextGenFishes[timer-1] += fishes
			}
		}
		currentFishes = nextGenFishes
	}

	var fishesTotal int
	for _, fishes := range currentFishes {
		fishesTotal += fishes
	}
	return fishesTotal
}
