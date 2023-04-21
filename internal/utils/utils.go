package utils

import (
	"bufio"
	"fmt"
	"math"
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
	if len(s) == 0 {
		return
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return
}

func Min(nums ...int) (min int) {
	min = math.MaxInt
	for _, num := range nums {
		if num < min {
			min = num
		}
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
