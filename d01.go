package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func doCheck(line string, c string, i int, start bool) int {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	ln := len(numbers)

	n, err := strconv.Atoi(c)
	if err == nil {
		return n
	} else {
		var lslice string
		if start {
			lslice = line[:i+1]
		} else {
			lslice = line[i:]
		}
		for j := 0; j < ln; j++ {
			if strings.Contains(lslice, numbers[j]) {
				return j + 1
			}
		}
	}
	return 0
}

func getNumber(line string) int {
	if len(line) == 0 {
		return 0
	}
	ll := len(line) - 1
	i := 0
	var first int
	var second int

	for first == 0 || second == 0 {
		if first == 0 {
			first = doCheck(line, string(line[i]), i, true)
		}
		if second == 0 {
			second = doCheck(line, string(line[ll-i]), ll-i, false)
		}
		i++
	}
	return (first * 10) + second
}

func sumD01Slice(n []int) int {
	result := 0
	for i := 0; i < len(n); i++ {
		result += n[i]
	}
	return result
}

func ReadLines(filePath string) (int, error) {
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
	return sumD01Slice(nums), nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Enter filename of test file")
	}

	filePath := os.Args[1]
	sum, err := ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
