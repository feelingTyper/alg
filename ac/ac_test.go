package ac

import (
	"fmt"
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
			Content:  "我是红领巾，祖国未来的花朵",
			Patterns: []string{"红领巾", "祖国", "花朵", "我的", "祖国未来", "国未"},
			Expects:  []string{"红领巾", "祖国", "国未", "祖国未来", "花朵"},
		},
		{
			Content:  "somi anymore",
			Patterns: []string{"anymore", "i am blessed", "i a"},
			Expects:  []string{"i a", "anymore"},
		},
	}

	for _, tc := range tests {
		ac := NewAcMachine()
		results := ac.AddPatterns(tc.Patterns...).Build().Print().SimpleQuery(tc.Content)
		assert.Equal(t, tc.Expects, results)

		fmt.Printf("patterns=%v, content=%s\n", tc.Patterns, tc.Content)
		for _, result := range results {
			fmt.Println(result)
		}
	}
}
