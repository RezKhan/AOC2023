package main

import (
	"log"
	rf "readfile"
	"slices"
	"strconv"
	"strings"
)

var cardfaces = map[string]int {
	"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7":7, 
	"8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
}

type hand struct {
	cards []string
	bid int
}


// sequences
// 5 of one
// 4 of one -> highest first card
// 3 of one, 2 of another -> highest first card
// 3 of one, 1 of one, 1 of one -> highest first card
// 2 of one, 2 of one, 1 of one -> highest first card
// Only ones -> highest first card



func partOne(lines []string) {
	var hands []hand

	for _, line := range lines {
		var h  hand
		tmp := strings.Split(line, " ")	
		cards := strings.Split(tmp[0], "")
		bid, err := strconv.Atoi(tmp[1]) 
		if err != nil {
			log.Println(err)
		}
		h.cards = cards
		h.bid = bid
		hands = append(hands, h)		
	}
	for _, h := range hands {
		log.Println(h)
		c := make([]string, len(h.cards))
		copy(c, h.cards)
		slices.Sort(c)
		log.Println(h.cards, c)
	}
}

func main()  {
	filePath := "d07test.txt"

	lines := rf.ReadFile(filePath)
	partOne(lines)
}
