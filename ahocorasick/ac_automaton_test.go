package ahocorasick

import (
	"testing"
)

func TestACAutomaton(t *testing.T) {
	testCases := []struct {
		keywords []string
		text     string
		ans      map[string][]int
	}{
		{
			[]string{
				"he",
				"she",
				"his",
				"hers",
			},
			"ahishers",
			map[string][]int{
				"his":  {1},
				"she":  {3},
				"he":   {4},
				"hers": {4},
			},
		},
		{
			[]string{
				"say",
				"she",
				"shr",
				"he",
				"her",
			},
			"sherhsay",
			map[string][]int{
				"she": {0},
				"he":  {1},
				"her": {1},
				"say": {5},
			},
		},
	}

	for _, tc := range testCases {
		ac := NewACAutomaton(tc.keywords)
		ans := ac.Search(tc.text)
		if len(ans) != len(tc.ans) {
			t.Fatalf("expected %v, but got %v", tc.ans, ans)
		}
		for k, v := range ans {
			if len(v) != len(tc.ans[k]) {
				t.Fatalf("expected %v, but got %v", tc.ans[k], v)
			}
			for i := range v {
				if v[i] != tc.ans[k][i] {
					t.Fatalf("expected %v, but got %v", tc.ans[k], v)
				}
			}
		}
	}
}
