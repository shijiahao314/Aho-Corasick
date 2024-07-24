package ahocorasick

import "fmt"

type Node struct {
	ID       int
	Val      byte
	Parent   *Node
	Children map[byte]*Node
	IsEnd    bool
	Count    []int
	Fail     *Node
}

func (n *Node) String() string {
	keys := make([]byte, 0)
	for k := range n.Children {
		keys = append(keys, k)
	}
	if n.IsEnd {
		return fmt.Sprintf("[<%d> %c, Count:%v, Parent: [<%d> %c]]", n.ID, n.Val, n.Count, n.Parent.ID, n.Parent.Val)
	}
	return fmt.Sprintf("[<%d> %c -> %c, Parent: [<%d> %c]]", n.ID, n.Val, keys, n.Parent.ID, n.Parent.Val)
}

type ACAutomaton struct {
	Root *Node
}

func NewACAutomaton(keywords []string) *ACAutomaton {
	acAutomaton := &ACAutomaton{
		Root: &Node{
			Val:      '@',
			Children: make(map[byte]*Node),
		},
	}
	acAutomaton.buildTrie(keywords)
	acAutomaton.buildFail()
	return acAutomaton
}

func (ac *ACAutomaton) buildTrie(keywords []string) {
	id := 1
	for _, keyword := range keywords {
		cur := ac.Root
		for i := 0; i < len(keyword); i++ {
			c := keyword[i]
			if cur.Children[c] == nil {
				cur.Children[c] = &Node{
					ID:       id,
					Val:      c,
					Parent:   cur,
					Children: make(map[byte]*Node),
				}
				id++
			}
			if i == len(keyword)-1 {
				cur.Children[c].IsEnd = true
				cur.Children[c].Count = append(cur.Children[c].Count, i+1)
			}
			cur = cur.Children[c]
		}
	}
}

func (ac ACAutomaton) PrintTrie() {
	q := make([]*Node, 0)
	q = append(q, ac.Root)
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			fmt.Printf("%s\n", cur)
			for _, v := range cur.Children {
				q = append(q, v)
			}
		}
	}
}

func (ac *ACAutomaton) buildFail() {
	// 1.root的fail指向自己
	ac.Root.Fail = ac.Root
	q := make([]*Node, 0)
	for _, v := range ac.Root.Children {
		q = append(q, v)
	}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			for _, v := range cur.Children {
				q = append(q, v)
			}
			// 2.当前节点的fail指向root
			cur.Fail = ac.Root
			for p := cur.Parent.Fail; ; p = p.Fail {
				flag := false
				for k, v := range p.Children {
					// 查找父节点fail指向的节点p
					// 若p的子节点中存在与当前节点相同的字符
					// 则当前节点的fail指向该子节点
					// 若p的子节点中不存在与当前节点相同的字符
					// 则继续寻找p的fail指向的节点，直至root节点
					if cur != v && cur.Val == v.Val {
						cur.Fail = p.Children[k]
						if cur.Fail.IsEnd {
							cur.IsEnd = true // TODO:这里存疑
							cur.Count = append(cur.Count, cur.Fail.Count...)
							flag = true
						}
						break
					}
				}
				if flag || p == ac.Root {
					break
				}
			}
		}
	}
}

func (ac *ACAutomaton) Search(text string) []string {
	ans := make([]string, 0)

	idx := 0

	for idx < len(text) {
		cur := ac.Root
		for ; idx < len(text); idx++ {
			if cur.Children[text[idx]] == nil {
				cur = cur.Fail
				if cur == ac.Root {
					break
				}
			}
			cur = cur.Children[text[idx]]
			if cur.IsEnd {
				for _, k := range cur.Count {
					ans = append(ans, string(text[idx+1-k:idx+1]))
				}
			}
		}
	}

	return ans
}
