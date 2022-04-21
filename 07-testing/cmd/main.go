package main

import (
	"fmt"
	"go-core/07-testing/pkg/sorter"
)

func main() {
	printSortedInts()
	printSortedStr()
}

func printSortedStr() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sorter.SortString(s)
	fmt.Println(s)
}

func printSortedInts() {
	i := []int{5, 2, 6, 3, 1, 4}
	sorter.SortInts(i)
	fmt.Println(i)
}
