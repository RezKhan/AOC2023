package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)


func getNumber(line string) int {
	if len(line) == 0 {
		return 0
	}
	ll := len(line) - 1
	i := 0
	var first int
	var second int

	for (first == 0 || second == 0) {
		if first == 0 {
			f, err := strconv.Atoi(string(line[i]))
			if err == nil {
				first = f
			}
		}
		if second == 0 {
			s, err := strconv.Atoi(string(line[ll-i]))
			if err == nil {
				second = s
			}
		}
		i++
	}
	return (first * 10) + second
}


func sumSlice(n []int) int {
	result := 0
	for i := 0; i < len(n); i++ {
		result += n[i]
	}
	return result
}

func ReadFile(filePath string) (int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	
	var nums []int
	for fScanner.Scan() {
		line := fScanner.Text()
		n := getNumber(line)
		nums = append(nums, n)
	}

	return sumSlice(nums), nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Enter filename of test file")
	}

	filePath := os.Args[1]
	sum, err := ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
