package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() int64 {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	stringData := strings.Split(string(data), "\n")
	parts := make([][]string, 0)
	temp := []string{}
	for _, line := range stringData {
		if line == "" {
			parts = append(parts, temp)
			temp = []string{}
			continue
		}
		temp = append(temp, line)
	}
	locationMap := parts[1:]
	seeds := strings.Split(strings.Split(parts[0][0], ":")[1], " ")[1:]
	var minInt int64 = math.MaxInt64
	for _, seedNumber := range seeds {
		seed, err := strconv.ParseInt(seedNumber, 10, 64)
		if err != nil {
			panic(err)
		}
		index := 0
		for index < len(locationMap) {
			for _, ranges := range locationMap[index][1:] {
				rangesList := strings.Split(ranges, " ")
				end, _ := strconv.ParseInt(rangesList[0], 10, 64)
				start, _ := strconv.ParseInt(rangesList[1], 10, 64)
				count, _ := strconv.ParseInt(rangesList[2], 10, 64)
				if start+count > seed && start <= seed {
					seed = end + seed - start
					break
				}
			}
			index++
		}
		if minInt > seed {
			minInt = seed
		}
	}
	return minInt
}

func Part2() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	stringData := strings.Split(string(data), "\n")
	parts := make([][]string, 0)
	temp := []string{}
	for _, line := range stringData {
		if line == "" {
			parts = append(parts, temp)
			temp = []string{}
			continue
		}
		temp = append(temp, line)
	}
	newLocationMap := make([][][]int, 0)
	for _, ranges := range parts[1:] {
		locationMap := make([][]int, 0)
		for _, rangesList := range ranges[1:] {
			rangesList := strings.Split(rangesList, " ")
			end, _ := strconv.Atoi(rangesList[0])
			start, _ := strconv.Atoi(rangesList[1])
			count, _ := strconv.Atoi(rangesList[2])
			locationMap = append(locationMap, []int{end, start, count})
		}
		newLocationMap = append(newLocationMap, locationMap)
	}
	seeds := strings.Split(strings.Split(parts[0][0], ":")[1], " ")[1:]
	packs := []int{}
	for i := 0; i < len(seeds)/2; i++ {
		start, _ := strconv.Atoi(seeds[2*i])
		end, _ := strconv.Atoi(seeds[2*i+1])
		for index := start; index < start+end; index++ {
			packs = append(packs, index)
		}
	}
	for index := 0; index < 1; index++ {
		for packIndex := 0; packIndex < len(packs); packIndex++ {
			seed := packs[packIndex]
			for _, ranges := range newLocationMap[index] {
				end := ranges[0]
				start := ranges[1]
				count := ranges[2]
				if start+count > seed && start <= seed {
					seed = end + seed - start
					break
				}
			}
			packs[packIndex] = seed
		}
	}
	min := math.MaxInt
	for _, x := range packs {
		if min > x {
			min = x
		}
	}
	return min
}

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}
