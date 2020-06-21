package main

import "fmt"

/*
Implement an autocomplete system. That is, given a query string s and a set of all possible query strings, return all strings in the set that have s as a prefix.

For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].

Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.
*/

func main() {
	fmt.Println(getSuggestions("foo", []string{"foobar", "fobar", "fooobar", "barfoo"}))
}

type node struct {
	val   rune
	nodes map[rune]*node
}

func getSuggestions(s string, words []string) []string {
	nodes := make(map[rune]*node)
	for _, w := range words {
		nodeMap := nodes
		for _, let := range w {
			if nd, ok := nodeMap[let]; ok {
				nodeMap = nd.nodes
				continue
			}
			nodeMap[let] = &node{
				val:   let,
				nodes: make(map[rune]*node),
			}
			nodeMap = nodeMap[let].nodes
		}
	}
	var complNode *node
	nodeMap := nodes
	for _, let := range s {
		if nd, ok := nodeMap[let]; ok {
			nodeMap = nd.nodes
			complNode = nd
			continue
		}
		return []string{}
	}
	var f func(n *node, s string)
	var result []string
	f = func(n *node, s string) {
		if len(n.nodes) == 0 {
			result = append(result, s)
			return
		}
		for let, nd := range n.nodes {
			f(nd, s+string(let))
		}
	}
	f(complNode, s)

	return result
}
