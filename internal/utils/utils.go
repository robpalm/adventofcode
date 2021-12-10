package utils

import (
	"strconv"
)

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
