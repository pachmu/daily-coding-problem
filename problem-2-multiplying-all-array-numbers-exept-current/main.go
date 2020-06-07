package main

import "fmt"

func main() {
	fmt.Println(makeSliceOfMult([]int64{1, 2, 3, 4, 5}))
}

func makeSliceOfMult(s []int64) []int64 {
	var direct, inverse []int64
	var x, y int64
	for i := range s {
		if i == 0 {
			x = s[i]
			y = s[len(s)-1]
		} else {
			x = s[i] * x
			y = s[len(s)-1-i] * y
		}
		direct = append(direct, x)
		inverse = append(inverse, y)
	}
	var resp []int64
	for i := range s {
		var z int64
		if i == 0 {
			z = inverse[len(s)-2]
		} else if i == len(s)-1 {
			z = direct[i-1]
		} else {
			z = direct[i-1] * inverse[len(s)-i-2]
		}
		resp = append(resp, z)
	}
	return resp
}
