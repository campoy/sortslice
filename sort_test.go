package sortslice

import (
	"sort"
	"testing"
)

func BenchmarkSortIntLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.IntSlice([]int{1, 5, 2, 6, 4, 2, 8, 3, 2, 6}))
	}
}

func BenchmarkSortIntBis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(Int([]int{1, 5, 2, 6, 4, 2, 8, 3, 2, 6}))
	}
}

func BenchmarkSortInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(Any([]int{1, 5, 2, 6, 4, 2, 8, 3, 2, 6}))
	}
}

type integer int

func BenchmarkSortGenInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(Any([]integer{1, 5, 2, 6, 4, 2, 8, 3, 2, 6}))
	}
}

type name string

func TestStringKind(t *testing.T) {
	ss := []string{"one", "two", "three", "four", "five", "six", "seven"}
	ns := []name{"one", "two", "three", "four", "five", "six", "seven"}
	sort.Sort(Any(ss))
	sort.Sort(Any(ns))
	if len(ss) != len(ns) {
		t.Errorf("bad length %v vs %v", ss, ns)
	}
	for i := range ss {
		if ss[i] != string(ns[i]) {
			t.Errorf("mismatched %v: %v vs %v", i, ss[i], ns[i])
		}
	}
}
