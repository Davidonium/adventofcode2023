package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/davidonium/adventofcode2023/util"
)

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Game struct {
	ID    int
	Draws []Draw
}

type Draw struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	fd, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		os.Exit(1)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	var games []Game

	for scanner.Scan() {
		t := scanner.Text()
		parts := strings.SplitN(t, ":", 2)
		rawGame := parts[0]
		rawDraws := parts[1]

		gameParts := strings.Split(rawGame, " ")
		gameId := util.ParseInt(gameParts[1])

		game := Game{
			ID: gameId,
		}

		draws := strings.Split(rawDraws, ";")

		for _, s := range draws {
			draw := Draw{}
			sets := strings.Split(s, ",")

			for _, s := range sets {
				parts := strings.SplitN(strings.TrimSpace(s), " ", 2)
				amount := util.ParseInt(parts[0])
				color := parts[1]

				switch color {
				case "red":
					draw.Red = amount
				case "green":
					draw.Green = amount
				case "blue":
					draw.Blue = amount
				}
			}

			game.Draws = append(game.Draws, draw)
		}

		games = append(games, game)
	}

	var powers []int
	for _, g := range games {
		maxes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, d := range g.Draws {
			if d.Red > maxes["red"] {
				maxes["red"] = d.Red
			}
			if d.Green > maxes["green"] {
				maxes["green"] = d.Green
			}
			if d.Blue > maxes["blue"] {
				maxes["blue"] = d.Blue
			}
		}

		powers = append(powers, maxes["red"]*maxes["green"]*maxes["blue"])
	}

	fmt.Printf("the sum of the powers is %d\n", util.SumSlice(powers))

	/* Part 1
	var possibleGames []Game

	gameloop:
		for _, g := range games {
			for _, d := range g.Draws {
				if d.Red > limits["red"] || d.Green > limits["green"] || d.Blue > limits["blue"] {
					continue gameloop
				}
			}

			possibleGames = append(possibleGames, g)
		}

		sum := 0
		for _, g := range possibleGames {
			sum += g.ID
		}

		fmt.Printf("The add up of the possible game IDs is %d\n", sum)
	*/
}
