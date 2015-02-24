package main

import (
	"fmt"
	"sort"
)

func checkio(data []int) float32 {
	sort.Ints(data)
	l := len(data)
	mid := l / 2
	if l%2 != 0 {
		return float32(data[mid])
	}
	return float32(data[mid-1]+data[mid]) / 2.0
}

func main() {
	fmt.Println(checkio([]int{1, 2, 3, 4, 5}) == 3)
	fmt.Println(checkio([]int{3, 1, 2, 5, 3}) == 3)
	fmt.Println(checkio([]int{1, 300, 2, 200, 1}) == 2)
	fmt.Println(checkio([]int{3, 6, 20, 99, 10, 15}) == 12.5)
}
