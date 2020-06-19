package main

import "fmt"

/*

A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.

Given the root to a binary tree, count the number of unival subtrees.

For example, the following tree has 5 unival subtrees:

  0
 / \
1   0
   / \
  1   0
 / \
1   1
*/

func main() {
	fmt.Println(countUnival(
		&node{
			val: 0,
			left: &node{
				val: 1,
			},
			right: &node{
				val: 1,
				left: &node{
					val: 0,
					left: &node{
						val: 1,
					},
					right: &node{
						val: 1,
					},
				},
				right: &node{
					val: 0,
				},
			},
		}))
}

type node struct {
	val   int
	left  *node
	right *node
}

func countUnival(n *node) (int, bool) {
	if n.right == nil && n.left == nil {
		return 1, true
	}
	var isUnivalLeft, isUnivalRight bool
	sum := 0
	if n.left != nil {
		var s int
		s, isUnivalLeft = countUnival(n.left)
		sum += s
	}
	if n.right != nil {
		var s int
		s, isUnivalRight = countUnival(n.right)
		sum += s
	}
	if isUnivalRight && isUnivalLeft && n.right.val == n.val && n.val == n.left.val {
		return sum + 1, true
	}

	return sum, false
}
