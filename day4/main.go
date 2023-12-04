package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/davidonium/adventofcode2023/util"
)

type Card struct {
	ID      int
	Matches int
	Points  int
}

func main() {
	fd, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		os.Exit(1)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	var cards []Card
	for scanner.Scan() {
		t := scanner.Text()

		card, numbers, _ := strings.Cut(t, ":")
		_, cardID, _ := strings.Cut(card, " ")

		winning, have, _ := strings.Cut(numbers, "|")

		winners := parseNumbers(winning)
		owned := parseNumbers(have)

		var matches []int
		for n := range owned {
			if _, ok := winners[n]; ok {
				matches = append(matches, n)
			}
		}

		points := int(math.Pow(2, float64(len(matches))-1))

		cards = append(cards, Card{
			ID:      util.ParseInt(cardID),
			Matches: len(matches),
			Points:  points,
		})

		fmt.Printf("card %s has %d points\n", cardID, int(points))
	}

	scratchCards := make([]int, len(cards))
	for i, c := range cards {
		// iterate the copies number and the original
		n := scratchCards[i] + 1

		for k := 0; k < n; k++ {
			for j := 1; j <= c.Matches; j++ {
				scratchCards[i+j]++
			}
		}
	}

	// scratchCards only tracks copies, add the length of the originals
	fmt.Printf("the total number of cards is %d\n", util.SumSlice(scratchCards)+len(cards))
}

func parseNumbers(s string) map[int]bool {
	split := strings.Fields(s)
	ints := make(map[int]bool, len(split))
	for _, w := range split {
		ints[util.ParseInt(w)] = true
	}
	return ints
}
