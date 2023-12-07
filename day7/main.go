package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/davidonium/adventofcode2023/util"
)

const (
	High = iota + 1
	Pair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var scores = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type Hand struct {
	Cards string
	Bid   int
	Type  int
}

func main() {
	fd, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		os.Exit(1)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	var hands []*Hand
	for scanner.Scan() {
		t := scanner.Text()

		cards, bid, _ := strings.Cut(t, " ")
		hands = append(hands, &Hand{
			Cards: cards,
			Bid:   util.ParseInt(bid),
		})
	}

	for _, h := range hands {
		coinc := map[rune]int{}

		for _, card := range h.Cards {
			coinc[card]++
		}

		handType := High
		for _, v := range coinc {
			if v == 5 {
				handType = FiveOfAKind
				break
			}

			if v == 4 {
				handType = FourOfAKind
				break
			}

			if v == 3 {
				if handType == Pair {
					handType = FullHouse
				} else {
					handType = ThreeOfAKind
				}
			}

			if v == 2 {
				if handType == ThreeOfAKind {
					handType = FullHouse
				} else if handType == Pair {
					handType = TwoPair
				} else {
					handType = Pair
				}
			}
		}

		h.Type = handType
	}

	sort.SliceStable(hands, func(i, j int) bool {
		ta := hands[i]
		tb := hands[j]

		if ta.Type == tb.Type {
			fmt.Printf("comparing two hands with same type %d - %s vs %s\n", ta.Type, ta.Cards, tb.Cards)
			rca := []rune(ta.Cards)
			rcb := []rune(tb.Cards)

			for i := 0; i < len(rca); i++ {
				scorea := scores[rca[i]]
				scoreb := scores[rcb[i]]

				fmt.Printf("comparing %s to %s\n", string(rca[i]), string(rcb[i]))

				if scorea != scoreb {
					fmt.Printf("scores are different: %d vs %d\n", scorea, scoreb)
					return scorea < scoreb
				}
			}

			return true
		}

		return ta.Type < tb.Type
	})

	winnings := 0
	for i, h := range hands {
		winnings += (i + 1) * h.Bid
		fmt.Printf("rank %d - hand %s - type %d - bid %d\n", i+1, h.Cards, h.Type, h.Bid)
	}

	fmt.Printf("total winnings: %d\n", winnings)
}
