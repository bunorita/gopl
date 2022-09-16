package treesort

import (
	"strconv"
	"strings"
)

type Tree struct {
	value       int
	left, right *Tree
}

func NewTree(values []int) *Tree {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	return root
}

// ex 7.3
func (t *Tree) String() string {
	values := appendValues([]int{}, t)
	sValues := []string{}
	for _, v := range values {
		sValues = append(sValues, strconv.Itoa(v))
	}
	return "{" + strings.Join(sValues, " ") + "}"
}

// Sort sorts values in place.
func Sort(values []int) {
	root := NewTree(values)
	appendValues(values[:0], root)
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		return &Tree{value: value}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
