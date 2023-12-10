package main

import (
	"fmt"
	"log"
	rf "readfile"
	"slices"
	"strconv"

	// "slices"
	"strings"
)

func raceRecord(time int, distance int) int {
	var beats int
	for held := 0; held < time; held++ {
		speed := held
		travel := (time - held) * speed
		// log.Println("held for", held, "moved for", time - held, "travelled:", travel)
		if travel > distance {
			beats++
		}
	}
	return beats
}

func partTwo(lines []string) {
	var times []int
	var distances []int
	for _, line := range lines {
		temp := strings.Split(line, " ")
		for _, t := range temp {
			n, err := strconv.Atoi(t)
			if err != nil {
				continue
			}
			if slices.Contains(temp, "Time:") {
				times = append(times, n)
			}
			if slices.Contains(temp, "Distance:") {
				distances = append(distances, n)
			}
		}
	}
	t, _ := strconv.Atoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(times)), ""), "[]"))
	d, _ := strconv.Atoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(distances)), ""), "[]"))

	n := raceRecord(t, d)
	log.Println(n)
}

func partOne(lines []string) {
	var times []int
	var distances []int
	for _, line := range lines {
		temp := strings.Split(line, " ")
		for _, t := range temp {
			n, err := strconv.Atoi(t)
			if err != nil {
				continue
			}
			if slices.Contains(temp, "Time:") {
				times = append(times, n)
			}
			if slices.Contains(temp, "Distance:") {
				distances = append(distances, n)
			}
		}
	}
	combo := 1
	for i := range times {
		n := raceRecord(times[i], distances[i])
		log.Println("Race:", i+1, "- ways to win:", n)
		combo *= n
	}
	log.Println("Combination is:", combo)
}

func main() {
	// filePath := "d06test.txt"
	filePath := "d06input.txt"
	lines := rf.ReadFile(filePath)
	partOne(lines)
	partTwo(lines)
}
