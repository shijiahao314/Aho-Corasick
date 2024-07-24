package ahocorasick

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	ID       int
	Val      byte
	Children map[byte]*TrieNode
	Output   []int
	Fail     *TrieNode
}

func (n *TrieNode) String() string {
	keys := make([]byte, 0)
	for k := range n.Children {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return fmt.Sprintf("[<%d> %c, Output:%v]", n.ID, n.Val, n.Output)
	}
	return fmt.Sprintf("[<%d> %c -> %c]", n.ID, n.Val, keys)
}

var TrieNodeID = 0

func NewTrieNode(val byte) *TrieNode {
	TrieNodeID++
	return &TrieNode{
		ID:       TrieNodeID,
		Val:      val,
		Children: make(map[byte]*TrieNode),
	}
}

type ACAutomaton struct {
	Root *TrieNode
}

func (ac ACAutomaton) String() string {
	var sb strings.Builder
	line := "====================\n"
	sb.WriteString(line)
	q := make([]*TrieNode, 0)
	q = append(q, ac.Root)
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			sb.WriteString(cur.String() + "\n")
			for _, v := range cur.Children {
				q = append(q, v)
			}
		}
	}
	sb.WriteString(line)
	return sb.String()
}

func NewACAutomaton(keywords []string) *ACAutomaton {
	acAutomaton := &ACAutomaton{
		Root: NewTrieNode('@'),
	}
	acAutomaton.buildTrie(keywords)
	acAutomaton.buildFailPointers()
	return acAutomaton
}

func (ac *ACAutomaton) buildTrie(words []string) {
	for _, word := range words {
		cur := ac.Root
		for i := 0; i < len(word); i++ {
			ch := word[i]
			if _, ok := cur.Children[ch]; !ok {
				cur.Children[ch] = NewTrieNode(ch)
			}
			cur = cur.Children[ch]
		}
		cur.Output = append(cur.Output, len(word))
	}
}

func (ac *ACAutomaton) buildFailPointers() {
	q := make([]*TrieNode, 0)
	for _, v := range ac.Root.Children {
		v.Fail = ac.Root
		q = append(q, v)
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for c, child := range cur.Children {
			failNode := cur.Fail
			for failNode != nil {
				if next, ok := failNode.Children[c]; ok {
					child.Fail = next
					break
				}
				failNode = failNode.Fail
			}
			if failNode == nil {
				child.Fail = ac.Root
			}
			child.Output = append(child.Output, child.Fail.Output...)
			q = append(q, child)
		}
	}
}

func (ac *ACAutomaton) Search(text string) map[string][]int {
	res := make(map[string][]int, 0)
	cur := ac.Root

	for i := 0; i < len(text); i++ {
		ch := text[i]
		for cur != nil && cur.Children[ch] == nil {
			cur = cur.Fail
		}
		if cur == nil {
			cur = ac.Root
			continue
		}
		// cur != nil
		cur = cur.Children[ch]
		if len(cur.Output) > 0 {
			for _, v := range cur.Output {
				word := text[i-v+1 : i+1]
				res[word] = append(res[word], i-v+1)
			}
		}
	}

	return res
}
