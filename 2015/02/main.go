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
func Part(processer func(l int, w int, h int) int) int {
	fileScanner, fileCloser := ReadFile("input.txt")
	defer fileCloser.Close()
	surfaceArea := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := strings.Split(line, "x")
		l, w, h := numbers[0], numbers[1], numbers[2]
		length, _ := strconv.Atoi(l)
		width, _ := strconv.Atoi(w)
		height, _ := strconv.Atoi(h)
		surfaceArea += processer(length, width, height)
	}
	return surfaceArea
}

func SurfaceArea(length, width, height int) int {
	return 2*length*width + 2*width*height + 2*height*length + min(length*width, width*height, height*length)
}

func WrappingArea(length, width, height int) int {
	return length*width*height + min(length+width, width+height, height+length)*2
}

func main() {
	fmt.Println(Part(SurfaceArea))
	fmt.Println(Part(WrappingArea))
}
