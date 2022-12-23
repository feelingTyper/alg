package ac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ACTest struct {
	Content  string
	Patterns []string
	Expects  []string
}

func TestAc(t *testing.T) {
	tests := []*ACTest{
		{
			Content:  "sherhsay",
			Patterns: []string{"she", "he", "her"},
			Expects:  []string{"she", "he", "her"},
		},
		{
			Content:  "我的最爱是tpu",
			Patterns: []string{"我的", "tp"},
			Expects:  []string{"我的", "tp"},
		},
		{
			Content:  "somi anymore",
			Patterns: []string{"anymore", "i am blessed", "i a"},
			Expects:  []string{"i a", "anymore"},
		},
		{
			Content:  `// 当前节点下不存在匹配成功的下一个节点时所指向的以该节点为结尾的路径的最长后缀的末尾节点, 即 path[fail]是path[cur]的最长后缀, 这样当匹配到cur节点时，path[cur]上的模式串已经被找到，当cur.child[char](char是主串当前待匹配的字符)不存在时就要跳到path[fail]继续匹配fail.children[char]进行快速定位,无需回退主串.`,
			Patterns: []string{"char", "跳到", "fail", "ch", "节点时", "天下", "child", "path[c"},
			Expects:  []string{"节点时", "fail", "path[c", "节点时", "path[c", "ch", "child", "ch", "char", "ch", "char", "跳到", "fail", "fail", "ch", "child", "ch", "char"},
		},
	}

	for _, tc := range tests {
		ac := NewAcMachine()
		results := ac.AddPatterns(tc.Patterns...).Build().Print().SimpleQuery(tc.Content)
		assert.Equal(t, tc.Expects, results)
	}
}
