package ac

import (
	"fmt"
)

type Node struct {
	idx      int
	fail     *Node
	lengths  []int // 当前节点包含模式串的长度
	children map[rune]*Node
}

type Result struct {
	Idx  int
	Len  int
	Word string
}

type AcMachine struct {
	num  int
	root *Node
}

func newNode(idx int) *Node {
	return &Node{
		idx:      idx,
		fail:     nil,
		children: make(map[rune]*Node),
	}
}

func NewAcMachine() *AcMachine {
	return &AcMachine{
		num:  0,
		root: newNode(0),
	}
}

func (r *AcMachine) AddPatterns(ps ...string) *AcMachine {
	for _, p := range ps {
		r.addPattern(p)
	}

	return r
}

func (r *AcMachine) Num() int {
	return r.num
}

func (r *AcMachine) newNode(parent *Node) *Node {
	r.num++
	return newNode(r.num)
}

func (r *AcMachine) addPattern(p string) {
	if r.root == nil {
		r.root = r.newNode(nil)
	}

	iter := r.root
	for _, char := range []rune(p) {
		if _, ok := iter.children[char]; !ok {
			iter.children[char] = r.newNode(iter)
		}
		iter = iter.children[char]
	}

	iter.lengths = append(iter.lengths, len([]rune(p)))
}

// Build 构造fail节点，当前char节点的fail节点是parent.fail.children[char]
func (r *AcMachine) Build() *AcMachine {
	queue := make([]*Node, 0)

	for _, node := range r.root.children {
		node.fail = r.root
		queue = append(queue, node)
	}

	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]

		for char, child := range parent.children {
			failed := parent.fail
			for ; failed != nil; failed = failed.fail {
				if f, ok := failed.children[char]; ok {
					child.fail = f
					for _, l := range f.lengths {
						dup := false
						for _, ll := range child.lengths {
							if ll == l {
								dup = true
								break
							}
						}
						if !dup {
							child.lengths = append(child.lengths, l)
						}
					}
					break
				}
			}
			if failed == nil {
				child.fail = r.root
			}
			queue = append(queue, child)
		}
	}

	return r
}

func (r *AcMachine) Print() *AcMachine {
	fmt.Println(r.Debug())
	return r
}

func (r *AcMachine) Debug() (debug string) {
	traveling(r.root, "",
		func(node *Node, path string) error {
			if len(path) > 0 {
				debug = fmt.Sprintf("%snode=%d, path=%s[%d], lengths=%v, fail=%d\n", debug, node.idx, path, node.idx, node.lengths, node.fail.idx)
			}
			return nil
		})

	return
}

func traveling(node *Node, path string, fn func(*Node, string) error) error {
	if node == nil {
		return nil
	}

	if err := fn(node, path); err != nil {
		return err
	}

	for c, child := range node.children {
		if err := traveling(child, fmt.Sprintf("%s->%c", path, c), fn); err != nil {
			return err
		}
	}

	return nil
}

func (r *AcMachine) Query(ss string) []*Result {
	var (
		rs      = []rune(ss)
		iter    = r.root
		results = make([]*Result, 0)
	)

	for i, c := range rs {
		_, ok := iter.children[c]

		for !ok && iter != r.root {
			iter = iter.fail
			_, ok = iter.children[c]
		}

		if child, ok := iter.children[c]; ok {
			for _, l := range child.lengths {
				s := i - l + 1
				results = append(results, &Result{
					Idx:  s,
					Len:  l,
					Word: string(rs[s : i+1]),
				})
			}
			iter = child
		}
	}

	return results
}

func (r *AcMachine) SimpleQuery(ss string) []string {
	results := make([]string, 0)
	for _, r := range r.Query(ss) {
		results = append(results, r.Word)
	}
	return results
}
