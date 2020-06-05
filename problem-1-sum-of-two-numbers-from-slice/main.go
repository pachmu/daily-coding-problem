package main

import "fmt"

func main() {
	fmt.Println(isSumInSlice(
		[]int{1, -8, 3, 5, 4, 9, 2},
		1,
	),
	)
}

/*
This problem was recently asked by Google.

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/

func isSumInSlice(slice []int, num int) bool {
	if len(slice) == 0 {
		return false
	}
	dev := make(map[int]struct{})
	for _, s := range slice {
		if _, ok := dev[s]; ok {
			return true
		} else {
			dev[num-s] = struct{}{}
		}
	}

	return false
}
