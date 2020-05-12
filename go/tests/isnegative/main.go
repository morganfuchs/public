package main

import (
	"fmt"

	"./student"
	"github.com/01-edu/public/go/lib"
)

func isNegative(n int) {
	if n < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func main() {
	table := append(
		lib.MultRandInt(),
		lib.MinInt,
		lib.MaxInt,
		0,
	)
	for _, arg := range table {
		lib.Challenge("IsNegative", student.IsNegative, isNegative, arg)
	}
}