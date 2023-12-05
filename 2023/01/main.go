package main

import (
	"bufio"
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

func CheckWordInSubstring(words map[string]int, line string, i int) (bool, int) {
	for k, v := range words {
		if strings.HasPrefix(line[i:], k) {
			return true, v
		}
	}
	return false, 0
}

func findFirstNumberInLine(line string, words map[string]int) int {
	for i := 0; i < len(line); i++ {
		if shouldReturn, returnValue := CheckWordInSubstring(words, line, i); shouldReturn {
			return returnValue
		}
	}
	return -1
}

func findLastNumberInLine(line string, words map[string]int) int {
	for i := len(line) - 1; i >= 0; i-- {
		if shouldReturn, returnValue := CheckWordInSubstring(words, line, i); shouldReturn {
			return returnValue
		}
	}
	return -1
}

func ProcessText(digits map[string]int, file string) int {
	fileScanner, fileCloser := ReadFile(file)
	defer fileCloser.Close()
	number := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		number += findFirstNumberInLine(line, digits)*10 + findLastNumberInLine(line, digits)
	}
	return number
}

func main() {
	digits := map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	}
	println(ProcessText(digits, "input.txt"))
	digitsWithWords := map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	println(ProcessText(digitsWithWords, "input.txt"))
}
