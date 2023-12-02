#!/usr/bin/env bash

set -euo pipefail


if [[ -z $1 ]] 
then
    echo "please provide a day number"
    exit
fi

DAY_DIR=day$1

mkdir $DAY_DIR

touch $DAY_DIR/input.txt

tee -a $DAY_DIR/main.go << END
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fd, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		os.Exit(1)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		t := scanner.Text()
        // TODO start coding!
	}
}
END
