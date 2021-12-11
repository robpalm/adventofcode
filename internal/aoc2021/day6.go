package aoc2021

import (
	"bufio"
	"fmt"
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
	var fishes []int

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		for _, s := range parts {
			fishes = append(fishes, utils.ToInt(s))
		}
		fmt.Println(fishes)
	}

	for i := 0; i < days; i++ {
		var fishCount = len(fishes)
		for i := 0; i < fishCount; i++ {
			if fishes[i] == 0 {
				fishes[i] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[i]--
			}
		}
	}
	return len(fishes)
}
