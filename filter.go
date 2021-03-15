package main

import (
	"fmt"
)

func filter(array []int, fn func(int) bool) []int {
	res := []int{}
	for i, v := range array {
		if (fn(v)) {
			res = append(res, array[i])
		}
	}
	return res
}

func mapFunc(array []int, fn func(int) int) []int {
	for i, v := range array {
		array[i] = fn(v)
	}
	return array
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res := filter(mapFunc(array, func(e int) int {
		return e + 10
	}), func(e int) bool {
		return (e % 5) < 2
	})
	fmt.Println(res)

}

