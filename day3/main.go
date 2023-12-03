package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/davidonium/adventofcode2023/util"
)

type PartNumber struct {
	Value         int
	Len           int
	X             int
	Y             int
	IsInSchematic bool
}

type gear struct {
	X int
	Y int
}

func main() {
	fd, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		os.Exit(1)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	var lines []string
	for scanner.Scan() {
		t := scanner.Text()
		lines = append(lines, t)
	}

	var partNumbers []*PartNumber
	for i, l := range lines {
		var parsingNum []rune
		startIdx := -1

		for j, c := range l {
			if unicode.IsDigit(c) {
				if len(parsingNum) == 0 {
					startIdx = j
				}
				parsingNum = append(parsingNum, c)
			}

			// if it's the end of the line or the current rune is not a digit, flush the buffer and build the part
			if len(parsingNum) > 0 && (j == len(l)-1 || !unicode.IsDigit(c)) {
				v, _ := strconv.Atoi(string(parsingNum))
				partNumbers = append(partNumbers, &PartNumber{
					Value: v,
					Len:   len(parsingNum),
					X:     startIdx,
					Y:     i,
				})
				parsingNum = nil
				startIdx = -1
			}
		}
	}

	gears := map[gear][]*PartNumber{}

	for _, pn := range partNumbers {
		startx := 0
		if pn.X > 0 {
			startx = pn.X - 1
		}

		// check row above number
		if pn.Y > 0 {
			line := []rune(lines[pn.Y-1])
			endx := pn.X + pn.Len
			lineend := len(line) - 1
			if endx > lineend {
				endx = lineend
			}

			for i := startx; i <= endx; i++ {
				if isGear(line[i]) {
					g := gear{X: i, Y: pn.Y - 1}
					gears[g] = append(gears[g], pn)
				}
			}
		}

		// check row below number
		if pn.Y < len(lines)-1 {
			line := []rune(lines[pn.Y+1])
			endx := pn.X + pn.Len
			lineend := len(line) - 1
			if endx > lineend {
				endx = lineend
			}

			for i := startx; i <= endx; i++ {
				if isGear(line[i]) {
					g := gear{X: i, Y: pn.Y + 1}
					gears[g] = append(gears[g], pn)
				}
			}
		}

		line := []rune(lines[pn.Y])

		// check left of the number
		pos := pn.X - 1
		if pos > 0 {
			if isGear(line[pos]) {
				g := gear{X: pos, Y: pn.Y}
				gears[g] = append(gears[g], pn)
			}
		}

		// check right of the number
		endpos := pn.X + pn.Len
		if endpos < len(line) {
			if isGear(line[endpos]) {
				g := gear{X: endpos, Y: pn.Y}
				gears[g] = append(gears[g], pn)
			}
		}
	}

	var ratios []int
	for g, parts := range gears {
		if len(parts) == 2 {
			fmt.Printf("gear: (y:%d,x:%d) numbers:", g.Y, g.X)
			ratio := 1
			for _, pn := range parts {
				ratio *= pn.Value
				fmt.Printf("%d, ", pn.Value)
			}

			fmt.Printf(" - ratio: %d\n", ratio)

			ratios = append(ratios, ratio)
		}
	}

	fmt.Printf("the sum of the gear ratios is %d\n", util.SumSlice(ratios))
}

func isGear(r rune) bool {
	return r == '*'
}
