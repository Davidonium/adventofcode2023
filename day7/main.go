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

var handNames = map[int]string{
	High: "high", 
	Pair: "pair", 
	TwoPair: "two pair", 
	ThreeOfAKind: "three of a kind", 
	FullHouse: "full house", 
	FourOfAKind: "four of a kind", 
	FiveOfAKind: "five of a kind", 
}

var scores = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 1,
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

		if n, ok := coinc['J']; ok {
			switch handType {
			case FourOfAKind:
				switch n {
				case 4:
					fallthrough
				case 1:
					handType = FiveOfAKind
				}
			case FullHouse:
				switch n {
				case 3:
					fallthrough
				case 2:
					handType = FiveOfAKind
				}
			case ThreeOfAKind:
				switch n {
				case 3:
					fallthrough
				case 1:
					handType = FourOfAKind
				}
			case TwoPair:
				switch n {
				// two pair with JJ can turn into four of a kind
				case 2:
					handType = FourOfAKind
				case 1:
					handType = FullHouse
				}
			case Pair:
				switch n {
				// if the pair is the JJ, combine with any other card
				case 2:
					fallthrough
				case 1:
					handType = ThreeOfAKind
				}
			case High:
				if n == 1 {
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
			rca := []rune(ta.Cards)
			rcb := []rune(tb.Cards)

			for i := 0; i < len(rca); i++ {
				scorea := scores[rca[i]]
				scoreb := scores[rcb[i]]

				if scorea != scoreb {
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
		fmt.Printf("rank %d - hand %s - %s - bid %d\n", i+1, h.Cards, handNames[h.Type], h.Bid)
	}

	fmt.Printf("total winnings: %d\n", winnings)
}
