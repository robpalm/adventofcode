package utils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

type DayFunc func(*bufio.Scanner) int

func Solve(dayFunc DayFunc, input string) {
	file, _ := os.Open("../../assets/" + input + ".txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := dayFunc(scanner)

	funcName := runtime.FuncForPC(reflect.ValueOf(dayFunc).Pointer()).Name()
	fmt.Printf("%s %20d\n", funcName, result)
}

func ToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}

func Min(i1 int, i2 int) (i int) {
	i = i1
	if i2 < i1 {
		i = i2
	}
	return
}

func Max(i1 int, i2 int) (i int) {
	i = i1
	if i2 > i1 {
		i = i2
	}
	return
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
