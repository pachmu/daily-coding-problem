package main

import (
	"fmt"
	"strconv"
)

/*### Problem 362

This problem was asked by Twitter.

A strobogrammatic number is a positive number that appears the same after being rotated `180` degrees.
For example, `16891` is strobogrammatic.

Create a program that finds all strobogrammatic numbers with N digits.*/
func strobogrammatic(numCount int) []string {
	sbgNumbers := []int{1, 6, 8, 9, 0}
	n := numCount % 2

	var numList []string
	var f func(str string)
	f = func(str string) {
		if len(str) == numCount {
			numList = append(numList, str)
			return
		}
		for _, x := range sbgNumbers {
			y := strconv.Itoa(x)
			if str == "" && n > 0 {
				if x == 6 || x == 9 {
					continue
				}
				f(y)
				continue
			}
			if len(str) == numCount-2 && x == 0 {
				continue
			}
			z := y
			if x == 6 {
				z = "9"
			}
			if x == 9 {
				z = "6"
			}
			f(y + str + z)
		}
	}
	f("")

	return numList
}

func main() {
	fmt.Println(strobogrammatic(7))
}
