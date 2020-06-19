package main

import "fmt"

/*
Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.

For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.
*/

func main() {
	fmt.Println(maxSum([]int{2, 4, 6, 2, 5}))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxSum(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	if len(slice) == 1 {
		return slice[0]
	}

	firstSum := slice[0]
	secondSum := max(firstSum, slice[1])

	for _, elem := range slice[2:] {
		lastSecondSum := secondSum
		secondSum = max(secondSum, firstSum+elem)
		firstSum = lastSecondSum
	}
	return max(firstSum, secondSum)
}
