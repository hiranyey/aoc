package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Part1(stringMatch string) int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	count := 0
	for {
		hash := GetMD5Hash(string(data) + fmt.Sprintf("%d", count))
		if strings.HasPrefix(hash, stringMatch) {
			return count
		}
		count++
	}
}

func main() {
	fmt.Println(Part1("00000"))
	fmt.Println(Part1("000000"))
}
