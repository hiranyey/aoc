package main

import (
	"fmt"
	"os"
)

func Part1() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	floor := 0
	for _, character := range data {
		if character == '(' {
			floor++
		} else if character == ')' {
			floor--
		}
	}
	return floor
}

func Part2() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	floor := 0
	for index, character := range data {
		if character == '(' {
			floor++
		} else if character == ')' {
			floor--
		}
		if floor == -1 {
			return index + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}
