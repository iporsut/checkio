package main

import (
	"fmt"
	"reflect"
)

func checkio(list []int) (r []int) {
	var m map[int]int = make(map[int]int)
	r = []int{}

	for _, l := range list {
		if _, ok := m[l]; ok {
			m[l]++
		} else {
			m[l] = 1
		}
	}

	for _, l := range list {
		if m[l] > 1 {
			r = append(r, l)
		}
	}
	return
}

func main() {
	fmt.Println(reflect.DeepEqual(checkio([]int{1, 2, 3, 1, 3}), []int{1, 3, 1, 3}))
	fmt.Println(reflect.DeepEqual(checkio([]int{1, 2, 3, 4, 5}), []int{}))
	fmt.Println(reflect.DeepEqual(checkio([]int{5, 5, 5, 5, 5}), []int{5, 5, 5, 5, 5}))
	fmt.Println(reflect.DeepEqual(checkio([]int{10, 9, 10, 10, 9, 8}), []int{10, 9, 10, 10, 9}))
}
