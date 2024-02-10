package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	times := strings.Fields((strings.Split(lines[0], ":")[1]))
	distance := strings.Fields(strings.Split(lines[1], ":")[1])
	product := 1
	for i := 0; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distance[i])
		count := 0
		for time := 0; time <= t; time++ {
			if time*(t-time)-d > 0 {
				count++
			}
		}
		product *= count
	}
	return product
}

func Part2() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	times := strings.Join(strings.Fields((strings.Split(lines[0], ":")[1])), "")
	distance := strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), "")
	finalTime, _ := strconv.Atoi(times)
	finalDistance, _ := strconv.Atoi(distance)
	firstRace := 0
	lastRace := 0
	for time := 0; time <= finalTime; time++ {
		if time*(finalTime-time)-finalDistance > 0 {
			firstRace = time
			break
		}
	}
	for time := finalTime; finalTime >= 0; time-- {
		if time*(finalTime-time)-finalDistance > 0 {
			lastRace = time
			break
		}
	}
	return lastRace - firstRace + 1
}

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}
