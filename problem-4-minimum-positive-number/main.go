package main

import "fmt"

//Given an array of integers, find the first missing positive integer in linear time and constant space.
//In other words, find the lowest positive integer that does not exist in the array. The array can contain
//duplicates and negative numbers as well.
//
//For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.
//
//You can modify the input array in-place.

func main() {
	fmt.Println(findMinPositiveNotInSlice(
		[]int{
			2,
			3,
			2,
			6,
			7,
			5,
			0,
			2,
			-1,
		},
	))
}

func findMinPositiveNotInSlice(s []int) int {
	if len(s) == 0 {
		return 1
	}
	for i := range s {
		for s[i] > 0 && s[i] < len(s) && s[i] != i {
			v := s[i]
			s[i], s[v] = s[v], s[i]
			if s[i] == v {
				break
			}
		}
	}
	for i := 1; i < len(s); i++ {
		if s[i] != i {
			return i
		}
	}

	return len(s)
}
