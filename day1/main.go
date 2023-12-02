package main

import (
	"bufio"
	"fmt"
	"os"

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

	var calibrations []int
	for scanner.Scan() {
		var digits []int
		t := scanner.Text()

		for i := 0; i < len(t); i++ {
			c := t[i]

			var n int
			isNumber := false
			if c > '0' && c <= '9' {
				n = util.ParseInt(string(c))
				isNumber = true
			} else {
				for nw, v := range numberWords {
					nwLen := len(nw)
					if i+nwLen > len(t) {
						continue
					}
					wordToCheck := t[i : i+nwLen]
					if wordToCheck == nw {
						n = v

						isNumber = true
						break
					}
				}
			}

			if isNumber {
				digits = append(digits, n)
			}
		}

		sum := digits[0]*10 + digits[len(digits)-1]

		calibrations = append(calibrations, sum)

		fmt.Printf("%s - %d / %d = %d\n", t, digits[0], digits[len(digits)-1], sum)
	}

	fmt.Printf("sum of calibrations is: %d\n", util.SumSlice(calibrations))
}

var numberWords = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
