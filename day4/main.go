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
	ID string
	Points int
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
			ID: cardID,
			Points: points,
		})

		fmt.Printf("card %s has %d points\n", cardID, int(points))
	}

	total := 0
	for _, c := range cards {
		total += c.Points	
	}

	fmt.Printf("the sum of winning points is %d\n", total)
}

func parseNumbers(s string) map[int]bool {
	split := strings.Fields(s)
	ints := make(map[int]bool, len(split))
	for _, w := range split {
		ints[util.ParseInt(w)] = true
	}
	return ints
}
