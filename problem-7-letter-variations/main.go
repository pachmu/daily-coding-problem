package main

import (
	"fmt"
	"strconv"
)

/*
Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.

For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.

You can assume that the messages are decodable. For example, '001' is not allowed.
*/

func main() {
	fmt.Println(getVariationsNum("11223344"))
}

func getVariationsNum(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	var cache = make([]int, len(s))
	cache[len(s)-1] = 1

	for i := len(s) - 2; i >= 0; i-- {
		k := s[i : i+2]
		atoi, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		if atoi < 26 {
			cache[i]++
		}
		cache[i] += cache[i+1] + 1
	}
	return cache[0]
}
