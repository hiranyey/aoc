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
	gifts := map[[2]int]int{
		{0, 0}: 1,
	}
	currentPosition := [2]int{0, 0}
	for _, x := range data {
		switch x {
		case '^':
			currentPosition[0]++
		case 'v':
			currentPosition[0]--
		case '>':
			currentPosition[1]++
		case '<':
			currentPosition[1]--
		}
		gifts[currentPosition]++
	}
	return len(gifts)
}

func Part2() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	gifts := map[[2]int]int{
		{0, 0}: 2,
	}
	santaPosition := [2]int{0, 0}
	robotPosition := [2]int{0, 0}
	for i, x := range data {
		if i%2 == 0 {
			switch x {
			case '^':
				santaPosition[0]++
			case 'v':
				santaPosition[0]--
			case '>':
				santaPosition[1]++
			case '<':
				santaPosition[1]--
			}
			gifts[santaPosition]++
		} else {
			switch x {
			case '^':
				robotPosition[0]++
			case 'v':
				robotPosition[0]--
			case '>':
				robotPosition[1]++
			case '<':
				robotPosition[1]--
			}
			gifts[robotPosition]++
		}
	}
	return len(gifts)
}

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}
