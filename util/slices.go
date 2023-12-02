package util

func SumSlice(sl []int) int {
	s := 0
	for _, v := range sl {
		s += v
	}

	return s
}
