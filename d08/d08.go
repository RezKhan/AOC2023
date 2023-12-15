package main

import (
	"log"
	rf "readfile"
	"strings"
)

func compareNodes(a, b map[string][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for keyA := range a {
		if _, exists := b[keyA]; !exists { // Check if keyA exists in map b
			return false
		}
	}
	return true
}

func PartOne(lines []string) {
	nodes := make(map[string][]string)
	start := "AAA"
	end := "ZZZ"
	// R := 1
	// L := 0
	reachedEnd := false
	directions := strings.Split(lines[0], "")

	for _, line := range lines {
		if n := strings.Index(line, "="); n > 0 {
			name := line[:n-1]
			paths := strings.Split(strings.Trim(line[n+2:], " ()"), ", ")
			nodes[name] = paths
		}
	}
	counter := 0
	currentNode := map[string][]string{start: nodes[start]}
	targetNode := map[string][]string{end: nodes[end]}
	for !reachedEnd {
		if compareNodes(currentNode, targetNode) {
			reachedEnd = true
		} else {
			mc := counter % len(directions)
			if directions[mc] == "L" {
				for _, value := range currentNode {
					currentNode = map[string][]string{value[0]: nodes[value[0]]}
				}
			}
			if directions[mc] == "R" {
				for _, value := range currentNode {
					currentNode = map[string][]string{value[1]: nodes[value[1]]}
				}
			}
			counter++
		}
	}
	log.Println(currentNode)
	log.Println(counter)
}

func main() {
	// filePath := "d08test.txt"
	filePath := "d08input.txt"
	lines := rf.ReadFile(filePath)

	PartOne(lines)
}
