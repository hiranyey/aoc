package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumberStruct struct {
	number int
	start  int
	end    int
}

func checkIfSymbolIsPresent(s byte) bool {
	return s != '.' && (s < '0' || s > '9')
}

func findEnginePartNumbers(lines []string, lineNumber int) int {
	currentLine := lines[lineNumber]
	numberList := []NumberStruct{}
	currentTemp := 0
	start := 0
	for i := 0; i < len(currentLine); i++ {
		if currentLine[i] >= '0' && currentLine[i] <= '9' {
			currentTemp = currentTemp*10 + int(currentLine[i]-'0')
		} else {
			if currentTemp != 0 {
				numberList = append(numberList, NumberStruct{
					number: currentTemp,
					start:  start,
					end:    i,
				})
			}
			start = i + 1
			currentTemp = 0
		}
	}
	if currentTemp != 0 {
		numberList = append(numberList, NumberStruct{
			number: currentTemp,
			start:  start,
			end:    len(currentLine),
		})
	}
	sum := 0
	for _, number := range numberList {
		foundSymbol := false
		if lineNumber > 0 {
			aboveLine := lines[lineNumber-1]
			for i := number.start - 1; i < number.end+1; i++ {
				if i >= 0 && i < len(aboveLine) && checkIfSymbolIsPresent(aboveLine[i]) {
					foundSymbol = true
				}
			}
		}
		for i := number.start - 1; i < number.end+1; i++ {
			if i >= 0 && i < len(currentLine) && checkIfSymbolIsPresent(currentLine[i]) {
				foundSymbol = true
			}
		}
		if lineNumber < len(lines)-1 {
			belowLine := lines[lineNumber+1]
			for i := number.start - 1; i < number.end+1; i++ {
				if i >= 0 && i < len(belowLine) && checkIfSymbolIsPresent(belowLine[i]) {
					foundSymbol = true
				}
			}
		}
		if foundSymbol {
			sum += number.number
		}
	}
	return sum
}

func Part1() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	answer := 0
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i++ {
		answer += findEnginePartNumbers(lines, i)
	}
	return answer
}

func findNumbersInLine(lines []string, lineNumber int, index int) []int {
	numbers := []int{}
	if lineNumber >= 0 || lineNumber < len(lines) {
		currentLine := lines[lineNumber]
		foundMidNumber := currentLine[index] >= '0' && currentLine[index] <= '9'
		midNumber, _ := strconv.Atoi(string(currentLine[index]))
		foundLeftNumber := false
		leftNumber := 0
		foundRightNumber := false
		rightNumber := 0
		if index > 0 && currentLine[index-1] >= '0' && currentLine[index-1] <= '9' {
			foundLeftNumber = true
			tempNumber := ""
			tempIndex := index - 1
			for tempIndex >= 0 && currentLine[tempIndex] >= '0' && currentLine[tempIndex] <= '9' {
				tempNumber = string(currentLine[tempIndex]) + tempNumber
				tempIndex--
			}
			leftNumber, _ = strconv.Atoi(tempNumber)
		}
		if index < len(currentLine) && currentLine[index+1] >= '0' && currentLine[index+1] <= '9' {
			foundRightNumber = true
			tempNumber := ""
			tempIndex := index + 1
			for tempIndex <= len(currentLine)-1 && currentLine[tempIndex] >= '0' && currentLine[tempIndex] <= '9' {
				tempNumber = tempNumber + string(currentLine[tempIndex])
				tempIndex++
			}
			rightNumber, _ = strconv.Atoi(tempNumber)
		}
		if foundMidNumber {
			if foundLeftNumber && foundRightNumber {
				specialNumber, _ := strconv.Atoi(fmt.Sprintf("%d%d%d", leftNumber, midNumber, rightNumber))
				numbers = append(numbers, specialNumber)
			} else if foundLeftNumber {
				specialNumber, _ := strconv.Atoi(fmt.Sprintf("%d%d", leftNumber, midNumber))
				numbers = append(numbers, specialNumber)
			} else if foundRightNumber {
				specialNumber, _ := strconv.Atoi(fmt.Sprintf("%d%d", midNumber, rightNumber))
				numbers = append(numbers, specialNumber)
			}
		} else {
			if foundLeftNumber {
				numbers = append(numbers, leftNumber)
			}
			if foundRightNumber {
				numbers = append(numbers, rightNumber)
			}
		}
		fmt.Println(foundMidNumber, midNumber, foundLeftNumber, leftNumber, foundRightNumber, rightNumber)
	}
	return numbers
}

func findGearNumbers(lines []string, lineNumber int) int {
	currentLine := lines[lineNumber]
	sum := 0
	for i := 0; i < len(currentLine); i++ {
		if currentLine[i] == '*' {
			numbers := []int{}
			numbers = append(numbers, findNumbersInLine(lines, lineNumber-1, i)...)
			numbers = append(numbers, findNumbersInLine(lines, lineNumber, i)...)
			numbers = append(numbers, findNumbersInLine(lines, lineNumber+1, i)...)
			fmt.Println(numbers)
			if len(numbers) == 2 {
				product := 1
				for _, v := range numbers {
					product *= v
				}
				sum += product
			}
		}
	}
	return sum
}

func Part2() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	answer := 0
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i++ {
		answer += findGearNumbers(lines, i)
	}
	return answer
}

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}
