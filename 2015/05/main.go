package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func CheckIfWordIsNaughty(line string) bool {
	numberOfVowels := 0
	foundDoubleCharacter := false
	foundMonotonousCharacter := false
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	tempChar := byte(0)
	for i := 0; i < len(line); i++ {
		if tempChar == line[i] {
			foundDoubleCharacter = true
		}
		if line[i]-tempChar == 1 && (tempChar == byte('a') || tempChar == byte('c') || tempChar == byte('p') || tempChar == byte('x')) {
			foundMonotonousCharacter = true
		}
		if slices.Contains(vowels, line[i]) {
			numberOfVowels++
		}
		tempChar = line[i]
	}
	return numberOfVowels > 2 && foundDoubleCharacter && !foundMonotonousCharacter
}

func CheckIfWordIsNaughtyPartTwo(line string) bool {
	valleyCheck := false
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			valleyCheck = true
		}
	}
	if !valleyCheck {
		return valleyCheck
	}
	for i := 0; i < len(line)-2; i++ {
		tempChar := line[i]
		tempChar2 := line[i+1]
		for j := i + 2; j < len(line)-1; j++ {
			if tempChar == line[j] && tempChar2 == line[j+1] {
				return true
			}
		}
	}
	return false
}

func Part(Processor func(line string) bool) int {
	fileScanner, fileCloser := ReadFile("input.txt")
	defer fileCloser.Close()
	number := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if Processor(line) {
			number++
		}
	}
	return number
}

func main() {
	fmt.Println(Part(CheckIfWordIsNaughty))
	fmt.Println(Part(CheckIfWordIsNaughtyPartTwo))
}
