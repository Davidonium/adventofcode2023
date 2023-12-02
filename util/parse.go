package util

import (
	"log"
	"strconv"
)

func ParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("failed to parse \"%s\" to an integer. Returned 0 for dev speed.", s)
		return 0
	}

	return v
}

func ParseUInt64(s string) uint64 {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Printf("failed to parse \"%s\" to an integer. Returned 0 for dev speed.", s)
		return 0
	}

	return v
}
