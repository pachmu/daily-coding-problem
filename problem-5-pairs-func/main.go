package main

import "fmt"

/*cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.

Given this implementation of cons:

def cons(a, b):
	def pair(f):
		return f(a, b)
	return pair
Implement car and cdr.*/

func cons(a, b string) func(f func(a, b string) string) string {
	return func(f func(a, b string) string) string {
		return f(a, b)
	}
}

func car(fn func(f func(a, b string) string) string) string {
	return fn(func(a, b string) string {
		return a
	})
}

func cdr(fn func(f func(a, b string) string) string) string {
	return fn(func(a, b string) string {
		return b
	})
}

func main() {
	fmt.Println(car(cons("1", "2")))
	fmt.Println(cdr(cons("1", "2")))
}
