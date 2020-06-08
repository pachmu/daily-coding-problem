package main

import (
	"encoding/base64"
	"fmt"
)

//Given the root to a binary tree, implement serialize(root), which serializes the tree into a string,
//and deserialize(s), which deserializes the string back into the tree.

type node struct {
	val   string
	left  *node
	right *node
}

func main() {
	n := node{
		val: "5",
		right: &node{
			val: "7",
			right: &node{
				val:   "9",
				right: nil,
				left:  nil, //&node{
				//	val:   "8",
				//	right: nil,
				//	left:  nil,
				//},
			},
			left: &node{
				val:   "6",
				right: nil,
				left:  nil,
			},
		},
		left: &node{
			val: "3",
			right: &node{
				val:   "4",
				right: nil,
				left:  nil,
			},
			left: &node{
				val:   "2",
				right: nil,
				left:  nil,
			},
		},
	}
	s := serialize(&n)
	fmt.Println(s)

	resp, _, err := deserialize(s, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func serialize(n *node) string {
	s := "("
	if n.left != nil {
		s += serialize(n.left)
	} else {
		s += "nil "
	}
	s += base64.StdEncoding.EncodeToString([]byte(n.val))
	if n.right != nil {
		s += serialize(n.right)
	} else {
		s += " nil"
	}
	s += ")"

	return s
}

func deserialize(tree string, curSymbol int) (*node, int, error) {
	nd := node{}
	var err error

	if tree[0] == '(' {
		nd.left, curSymbol, err = deserialize(tree[1:], curSymbol+1)
		if err != nil {
			return nil, 0, err
		}
	} else {
		curSymbol = len("nil ")
	}

	tree = tree[curSymbol:]
	for i, s := range tree {
		var curString string
		if s == '(' {
			curString = tree[0:i]
			var additionalSymbol int
			nd.right, additionalSymbol, err = deserialize(tree[i+1:], curSymbol+i+1)
			if err != nil {
				return nil, 0, err
			}
			curSymbol += additionalSymbol + len(curString)

		} else if len(tree) >= i+4 && tree[i:i+4] == " nil" {
			curString = tree[0:i]
			curSymbol += i + 4
		} else {
			continue
		}
		e, err := base64.StdEncoding.DecodeString(curString)
		if err != nil {
			return nil, 0, err
		}
		nd.val = string(e)
		curSymbol += 2
		break
	}
	return &nd, curSymbol, nil

}
