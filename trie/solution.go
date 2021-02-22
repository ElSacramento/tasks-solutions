package trie

type Node struct {
	value          rune
	word           string
	isCompleteWord bool
	children       map[rune]*Node
}

func CreateTrie(words []string) *Node {
	root := &Node{children: make(map[rune]*Node, 26)}
	// O(n * m)
	for _, w := range words {
		current := root
		for _, s := range w {
			node, found := current.children[s]
			if found {
				// just go to the next character
				current = node
				continue
			}
			// not found: add new node, update children
			newNode := &Node{children: make(map[rune]*Node, 26), value: s, word: current.word + string(s)}
			current.children[s] = newNode
			current = newNode
		}
		current.isCompleteWord = true
	}
	return root
}

func isLower(elem rune) bool {
	return elem >= 97 && elem <= 122
}

func goDeep(node *Node, founded map[string]struct{}) {
	if node.isCompleteWord {
		founded[node.word] = struct{}{}
	}
	for value, ref := range node.children {
		if !isLower(value) {
			return
		}
		goDeep(ref, founded)
	}
}

func depthSearch(node *Node, pattern string, index int, founded map[string]struct{}) {
	if index >= len(pattern) {
		goDeep(node, founded)
		return
	}

	elem := rune(pattern[index])
	for value, ref := range node.children {
		if elem == value {
			depthSearch(ref, pattern, index+1, founded)
			continue
		}
		// elem != value
		// we can ignore lowercase elements
		if isLower(value) {
			depthSearch(ref, pattern, index, founded)
		}
	}
}

func patternMatch(root *Node, pattern string) map[string]struct{} {
	foundWords := make(map[string]struct{}, 0)
	depthSearch(root, pattern, 0, foundWords)
	return foundWords
}

// check if queries are matching the pattern
// query matches the pattern if lowercase letters can be inserted between uppercase letters
// FooBar matches FB, but FooBarCar doesn't
// trie = prefix tree, nodes is characters
// leetcode: 1023
func camelMatch(queries []string, pattern string) []bool {
	root := CreateTrie(queries)

	matchingQueries := patternMatch(root, pattern)
	answers := make([]bool, 0, len(queries))
	// O(n)
	for _, q := range queries {
		if _, found := matchingQueries[q]; found {
			answers = append(answers, true)
			continue
		}
		answers = append(answers, false)
	}
	return answers
}
