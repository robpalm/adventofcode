package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"robpalm/adventofcode/internal/aoc2021"
	"runtime"
)

func main() {
	solve(aoc2021.Day1_1, "aoc2021/day1")
	solve(aoc2021.Day1_2, "aoc2021/day1")

	solve(aoc2021.Day2_1, "aoc2021/day2")
	solve(aoc2021.Day2_2, "aoc2021/day2")

	solve(aoc2021.Day3_1, "aoc2021/day3")
	solve(aoc2021.Day3_2, "aoc2021/day3")

	solve(aoc2021.Day4_1, "aoc2021/day4")
	solve(aoc2021.Day4_2, "aoc2021/day4")

	solve(aoc2021.Day5_1, "aoc2021/day5")
	solve(aoc2021.Day5_2, "aoc2021/day5")
}

type DayFunc func(*bufio.Scanner) int

func solve(dayFunc DayFunc, input string) {
	file, _ := os.Open("../../assets/" + input + ".txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := dayFunc(scanner)

	funcName := runtime.FuncForPC(reflect.ValueOf(dayFunc).Pointer()).Name()
	fmt.Printf("%s %12d\n", funcName, result)
}
