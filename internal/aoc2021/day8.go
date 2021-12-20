package aoc2021

import (
	"bufio"
	"sort"
	"strings"
)

func Day8_1(scanner *bufio.Scanner) (result int) {
	digits := make(map[int]int)

	for scanner.Scan() {
		output := false
		for _, word := range strings.Split(scanner.Text(), " ") {
			if word == "|" {
				output = true
				continue
			}
			if output == true {
				digits[len(word)]++
			}
		}
	}
	result = digits[2] + digits[4] + digits[3] + digits[7]
	return
}

func Day8_2(scanner *bufio.Scanner) (result int) {
	uniqueTexts := make(map[string]int)
	uniqueDigits := make(map[int]string)

	for scanner.Scan() {
		var outputDigits []string
		onOutput := false
		for _, word := range strings.Split(scanner.Text(), " ") {
			if word == "|" {
				onOutput = true
			} else {
				text := sortString(word)
				if onOutput {
					outputDigits = append(outputDigits, text)
				}
				uniqueTexts[text]++
				uniqueDigits[len(text)] = text
			}
		}

		textDigit := make(map[string]int)
		for uniqueText := range uniqueTexts {
			digit := -1
			switch len(uniqueText) {
			case 2:
				digit = 1
			case 3:
				digit = 7
			case 4:
				digit = 4
			case 7:
				digit = 8
			case 5:
				if !hasUniqueRune(uniqueDigits[2], uniqueText) {
					digit = 3
				} else if !hasUniqueRune(getUniqueRunes(uniqueDigits[4], uniqueDigits[2]), uniqueText) {
					digit = 5
				} else {
					digit = 2
				}
			case 6:
				if !hasUniqueRune(uniqueDigits[4], uniqueText) {
					digit = 9
				} else if !hasUniqueRune(uniqueDigits[2], uniqueText) {
					digit = 0
				} else {
					digit = 6
				}
			default:
				continue
			}
			// digitText[digit] = uniqueText
			textDigit[uniqueText] = digit
		}

		output := 0
		for _, outputDigit := range outputDigits {
			digit := textDigit[outputDigit]
			output = output*10 + digit
		}
		result += output
	}
	return
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func hasUniqueRune(src, dst string) bool {
	for _, r1 := range src {
		if !strings.ContainsRune(dst, r1) {
			return true
		}
	}
	return false
}
func getUniqueRunes(src, dst string) (u string) {
	for _, r1 := range src {
		if !strings.ContainsRune(dst, r1) {
			u += string(r1)
		}
	}
	return
}
