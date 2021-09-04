package main

import (
	"fmt"
)

func main() {
	/* creating map */
	numberMap := map[string]int{}
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println("numberMap :", numberMap)

	numberMap2 := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	fmt.Println("numberMap2 :", numberMap2)

	numberMap3 := make(map[string]int, 3) // 용량 생략 가능
	numberMap3["one"] = 1
	numberMap3["two"] = 2
	numberMap3["three"] = 3
	fmt.Println("numberMap3 :", numberMap3)

	for k, v := range numberMap {
		fmt.Println(k, v)
	}
	fmt.Println("----------------------")

	/* creating array */
	group1 := []int32{1, 4, 6}
	group2 := []int32{2, 4, 5}
	group3 := []int32{4, 6, 7}

	fmt.Println(group1)
	fmt.Println(group2)
	fmt.Println(group3)
	fmt.Println("----------------------")

	/* creating map by array */
	groupMap := make(map[string]string)

	groupMap[string(group1)] = "first"
	groupMap[string(group2)] = "second"
	groupMap[string(group3)] = "third"

	fmt.Println(groupMap)
}
