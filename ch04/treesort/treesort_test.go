package treesort_test

import (
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
