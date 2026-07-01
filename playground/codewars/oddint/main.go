package main

func FindOdd(seq []int) int {
	m := make(map[int]int)

	for _, v := range seq {
		m[v]++
	}

	for n, count := range m {
		if count%2 != 0 {
			return n
		}
	}
	return 0
}
