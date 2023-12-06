package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/davidonium/adventofcode2023/util"
)

func main() {
	fd, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		os.Exit(1)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	var times []uint64
	var distances []uint64

	for scanner.Scan() {
		t := scanner.Text()

		field, values, _ := strings.Cut(t, ":")

		if field == "Time" {
			times = append(times, util.ParseUInt64(strings.ReplaceAll(values, " ", "")) ) 
		} else if field == "Distance" {
			distances = append(distances, util.ParseUInt64(strings.ReplaceAll(values, " ", "")) ) 
		}
	}


	var ways []uint64
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]

		min := uint64(1)
		max := time - 1

		var beaters []uint64
		for j := min; j <= max; j++ {
			rem := time - j
			if rem * j > distance {
				beaters = append(beaters, j)
			}
		}

		ways = append(ways, uint64(len(beaters)))
	}

	final := uint64(1)
	for _, w := range ways {
		final *= w
	}

	fmt.Printf("times: %+v\n", times)
	fmt.Printf("distances: %+v\n", distances)
	fmt.Printf("ways: %+v\n", ways)
	fmt.Printf("final: %d\n", final)
}
