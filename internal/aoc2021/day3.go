package aoc2021

import (
	"bufio"
	"strconv"
)

func Day3_1(scanner *bufio.Scanner) int {
	var arr []int
	var rows, gamma, epsilon int

	for scanner.Scan() {
		rows++
		line := scanner.Text()
		if arr == nil {
			arr = make([]int, len(line))
		}
		for pos, bit := range line {
			if bit == '1' {
				arr[pos]++
			}
		}
	}

	for pos := 0; pos < len(arr); pos++ {
		gamma <<= 1
		if arr[pos] > rows/2 {
			gamma |= 1
		}
	}
	epsilon = ^gamma & (1<<len(arr) - 1)
	return gamma * epsilon
}

func Day3_2(scanner *bufio.Scanner) int {
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	var oxygene = get(&rows, '1')
	var co2 = get(&rows, '0')

	return int(oxygene) * int(co2)
}

func get(rows *[]string, value byte) (result int64) {
	var rowFilter []int
	var rowLength = len((*rows)[0])
	var ones, zeroes []int

	for pos := 0; pos < rowLength; pos++ {
		ones, zeroes = filterArray(rows, &rowFilter, pos)
		if value == '1' {
			if len(ones) < len(zeroes) {
				rowFilter = zeroes
			} else {
				rowFilter = ones
			}
		}
		if value == '0' {
			if len(ones) < len(zeroes) {
				rowFilter = ones
			} else {
				rowFilter = zeroes
			}
		}
		if len(rowFilter) == 1 {
			s := (*rows)[rowFilter[0]]
			result, _ = strconv.ParseInt(s, 2, 64)
			break
		}
	}
	return
}

func filterArray(rows *[]string, rowFilter *[]int, rowPos int) (ones []int, zeroes []int) {
	if *rowFilter == nil {
		for index, row := range *rows {
			if row[rowPos] == '1' {
				ones = append(ones, index)
			} else {
				zeroes = append(zeroes, index)
			}
		}
	} else {
		for _, val := range *rowFilter {
			row := (*rows)[val]
			if row[rowPos] == '1' {
				ones = append(ones, val)
			} else {
				zeroes = append(zeroes, val)
			}
		}
	}

	return
}
