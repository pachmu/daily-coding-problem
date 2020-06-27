package main

import "fmt"

func main() {
	fmt.Println(getSliceMaxVal([]int{10, 5, 2, 7, 8, 7}, 3))
	fmt.Println(getSliceMaxVal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
	fmt.Println(getSliceMaxVal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(getSliceMaxVal([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 3))
	fmt.Println(getSliceMaxVal([]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}, 3))
}

func getSliceMaxVal(s []int, k int) []int {
	if len(s) < k {
		return []int{}
	}
	var result []int
	var first *int
	max := &s[0]
	var secondMax *int
	for i := range s {
		if secondMax == nil && s[i] < *max {
			secondMax = &s[i]
		}

		if s[i] > *max {
			secondMax = max
			max = &s[i]
		} else if secondMax != nil && s[i] > *secondMax {
			secondMax = &s[i]
		}

		if i < k-1 {
			continue
		}
		if max == first {
			max = secondMax
		}
		result = append(result, *max)

		first = &s[i-(k-1)]
	}

	return result
}
