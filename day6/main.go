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

	var times []int
	var distances []int

	for scanner.Scan() {
		t := scanner.Text()

		f := strings.Fields(t)

		if f[0] == "Time:" {
			for _, s := range f[1:] {
				times = append(times, util.ParseInt(s))
			}
		} else if f[0] == "Distance:" {
			for _, s := range f[1:] {
				distances = append(distances, util.ParseInt(s))
			}
		}
	}


	var ways []int
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]

		min := 1
		max := time - 1

		var beaters []int
		for j := min; j <= max; j++ {
			rem := time - j
			if rem * j > distance {
				beaters = append(beaters, j)
			}
		}

		ways = append(ways, len(beaters))
	}

	final := 1
	for _, w := range ways {
		final *= w
	}

	fmt.Printf("times: %+v\n", times)
	fmt.Printf("distances: %+v\n", distances)
	fmt.Printf("ways: %+v\n", ways)
	fmt.Printf("final: %d\n", final)
}
