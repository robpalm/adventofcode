package aoc2021

import (
	"bufio"
	"robpalm/adventofcode/internal/utils"
	"strings"
)

type Line struct {
	x1, y1, x2, y2 int
	xDelta, yDelta int
	xStep, yStep   int
	straight       bool
}

func Day5_1(scanner *bufio.Scanner) int {
	return Day5(scanner, true)
}

func Day5_2(scanner *bufio.Scanner) int {
	return Day5(scanner, false)
}

func Day5(scanner *bufio.Scanner, onlyStraightLines bool) int {
	var lines []Line

	for scanner.Scan() {
		row := strings.ReplaceAll(scanner.Text(), " ", "")
		row = strings.ReplaceAll(row, "->", ",")
		parts := strings.Split(row, ",")

		lines = append(lines, parseLine(parts))
	}

	points := make(map[int]byte)

	for _, line := range lines {
		if onlyStraightLines && !line.straight {
			continue
		}
		if line.xDelta > line.yDelta {
			for i := 0; i <= line.xDelta; i++ {
				x := line.x1 + line.xStep*i
				y := line.y1 + line.yStep*i
				points[x<<16+y]++
			}
		} else {
			for i := 0; i <= line.yDelta; i++ {
				x := line.x1 + line.xStep*i
				y := line.y1 + line.yStep*i
				points[x<<16+y]++
			}
		}
	}
	var crosses int
	for _, count := range points {
		if count > 1 {
			crosses++
		}
	}

	return crosses
}

func parseLine(parts []string) (line Line) {
	line.x1 = utils.ToInt(parts[0])
	line.y1 = utils.ToInt(parts[1])
	line.x2 = utils.ToInt(parts[2])
	line.y2 = utils.ToInt(parts[3])
	line.straight = true

	if line.x1 < line.x2 {
		line.xStep = 1
	} else if line.x2 < line.x1 {
		line.xStep = -1
	}

	if line.y1 < line.y2 {
		line.yStep = 1
	} else if line.y2 < line.y1 {
		line.yStep = -1
	}

	line.xDelta = utils.Max(line.x1, line.x2) - utils.Min(line.x1, line.x2)
	line.yDelta = utils.Max(line.y1, line.y2) - utils.Min(line.y1, line.y2)

	if line.xDelta == line.yDelta {
		line.straight = false
	}

	return
}
