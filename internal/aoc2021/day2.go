package aoc2021

import (
	"bufio"
	"robpalm/adventofcode/internal/utils"
	"strings"
)

func Day2_1(scanner *bufio.Scanner) int {
	var depth, length int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		command := parts[0]
		units := utils.ToInt(parts[1])

		switch command {
		case "forward":
			length += units
		case "up":
			depth -= units
		case "down":
			depth += units
		}
	}
	return depth * length
}

func Day2_2(scanner *bufio.Scanner) int {
	var depth, length, aim int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		command := parts[0]
		units := utils.ToInt(parts[1])

		switch command {
		case "forward":
			length += units
			depth += aim * units
		case "up":
			aim -= units
		case "down":
			aim += units
		}
	}
	return depth * length
}
