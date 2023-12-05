package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/davidonium/adventofcode2023/util"
)

type TransformTable struct {
	SourceStart int
	DestStart   int
	Range       int
}

type Node struct {
	Name            string
	TransformTables []TransformTable
	Next            *Node
}

func (n *Node) FindLeave(s int) int {
	cur := n

	d := s
	for cur != nil {
		for _, t := range cur.TransformTables {
			if d >= t.SourceStart && d < t.SourceStart+t.Range {
				tmp := t.DestStart + (d - t.SourceStart)
				d = tmp
				break
			}
		}

		cur = cur.Next
	}

	return d
}

func main() {
	fd, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		os.Exit(1)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	var seeds []int

	rootNode := &Node{}

	curNode := rootNode

	for scanner.Scan() {
		t := scanner.Text()

		if strings.HasPrefix(t, "seeds:") {
			rawSeeds := strings.TrimPrefix(t, "seeds:")
			seeds = StringToIntSlice(rawSeeds)
		} else if strings.Contains(t, "map:") {
			curNode.Name = strings.TrimSuffix(t, " map:")
		} else if t != "" {
			rawTable := StringToIntSlice(t)
			table := TransformTable{
				DestStart:   rawTable[0],
				SourceStart: rawTable[1],
				Range:       rawTable[2],
			}
			curNode.TransformTables = append(curNode.TransformTables, table)
		} else {
			// the root node won't have a name when reaching the first empty line
			if curNode.Name != "" {
				curNode.Next = &Node{}
				curNode = curNode.Next
			}
		}
	}

	var locations []int

	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			l := rootNode.FindLeave(j)
			locations = append(locations, l)
		}
	}
	m := slices.Min(locations)

	fmt.Printf("min is %d\n", m)
}

func StringToIntSlice(s string) []int {
	var ints []int
	for _, s := range strings.Fields(s) {
		ints = append(ints, util.ParseInt(s))
	}

	return ints
}
