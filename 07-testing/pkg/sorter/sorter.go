package sorter

import "sort"

func SortString(s []string) []string {
	sort.Strings(s)
	return s
}

func SortInts(i []int) []int {
	sort.Ints(i)
	return i
}

func SortFloat64s(f []float64) []float64 {
	sort.Float64s(f)
	return f
}
