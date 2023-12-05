package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func CheckIfGameIsValid(line string, ballValidity map[string]int64) (bool, int) {
	parts := strings.Split(line, ":")
	gameNumber, err := strconv.ParseInt(strings.Split(parts[0], " ")[1], 10, 64)
	if err != nil {
		panic(err)
	}
	balls := strings.Split(parts[1], ";")
	for _, ball := range balls {
		ballParts := strings.Split(ball, ",")
		for _, ballPart := range ballParts {
			gamePart := strings.Split(strings.TrimSpace(ballPart), " ")
			numberOfBall, err := strconv.ParseInt(gamePart[0], 10, 64)
			if err != nil {
				panic(err)
			}
			if ballValidity[gamePart[1]] < numberOfBall {
				return false, 0
			}
		}
	}
	return true, int(gameNumber)
}

func CheckMinimumNumberOfBalls(line string) int {
	ballsResult := map[string]int64{
		"blue":  0,
		"red":   0,
		"green": 0,
	}
	parts := strings.Split(line, ":")
	balls := strings.Split(parts[1], ";")
	for _, ball := range balls {
		ballParts := strings.Split(ball, ",")
		for _, ballPart := range ballParts {
			gamePart := strings.Split(strings.TrimSpace(ballPart), " ")
			numberOfBall, err := strconv.ParseInt(gamePart[0], 10, 64)
			if err != nil {
				panic(err)
			}
			if ballsResult[gamePart[1]] < numberOfBall {
				ballsResult[gamePart[1]] = numberOfBall
			}
		}
	}
	return int(ballsResult["blue"] * ballsResult["red"] * ballsResult["green"])
}

func Part1() int {
	fileScanner, fileCloser := ReadFile("input.txt")
	defer fileCloser.Close()
	ballValidity := map[string]int64{
		"blue":  14,
		"red":   12,
		"green": 13,
	}
	number := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if valid, gameNumber := CheckIfGameIsValid(line, ballValidity); valid {
			number += gameNumber
		}
	}
	return number
}

func Part2() int {
	fileScanner, fileCloser := ReadFile("input.txt")
	defer fileCloser.Close()
	number := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		number += CheckMinimumNumberOfBalls(line)
	}
	return number
}

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}
