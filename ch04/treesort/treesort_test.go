package treesort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/bunorita/gopl/ch04/treesort"
)

func TestSort(t *testing.T) {
	t.Parallel()

	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50 // 0,1,..,49
	}

	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v\n", data)
	}
}

// ex7.3
func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		values []int
	}{
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{[]int{8, 9, 6, 7, 4, 5, 2, 3, 0, 1}},
		{[]int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	want := "{0 1 2 3 4 5 6 7 8 9}"

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("values=%v", tt.values), func(t *testing.T) {
			t.Parallel()

			tr := treesort.NewTree(tt.values)
			if got := tr.String(); got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})

	}
}
