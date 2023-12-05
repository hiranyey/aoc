package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func ReadFile(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner, file
}

func Part1() int {
	fileScanner, fileCloser := ReadFile("input.txt")
	defer fileCloser.Close()
	number := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := strings.Split(numbers[0], " ")
		winningNumberSet := map[string]bool{}
		for _, winningNumber := range winningNumbers {
			if len(winningNumber) != 0 {
				winningNumberSet[winningNumber] = true
			}
		}
		givenNumbers := strings.Split(numbers[1], " ")
		exponent := 0
		for _, givenNumber := range givenNumbers {
			if winningNumberSet[givenNumber] {
				exponent++
			}
		}
		number += int(math.Pow(2, float64(exponent-1)))
	}
	return number
}

func Part2() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	amountOfScratchCards := []int{}
	for i := 0; i < len(lines); i++ {
		amountOfScratchCards = append(amountOfScratchCards, 1)
	}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := strings.Split(numbers[0], " ")
		winningNumberSet := map[string]bool{}
		for _, winningNumber := range winningNumbers {
			if len(winningNumber) != 0 {
				winningNumberSet[winningNumber] = true
			}
		}
		givenNumbers := strings.Split(numbers[1], " ")
		numberOfMatchedGuard := 1
		for _, givenNumber := range givenNumbers {
			if winningNumberSet[givenNumber] && i+numberOfMatchedGuard < len(lines) {
				amountOfScratchCards[i+numberOfMatchedGuard] += amountOfScratchCards[i]
				numberOfMatchedGuard++
			}
		}
	}
	sumOfScratchCards := 0
	for i := 0; i < len(lines); i++ {
		sumOfScratchCards += amountOfScratchCards[i]
	}
	return sumOfScratchCards
}

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}
