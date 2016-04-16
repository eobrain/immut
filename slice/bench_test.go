package slice_test

import (
	"github.com/eobrain/immut"
	"github.com/eobrain/immut/slice"
	"testing"
)

var seq = slice.New()

/////////////////////////////////////////////////////////////////////////////

func init() {
	for i := 0; i < 100; i++ {
		seq = seq.AddFront(i)
	}
}

func BenchmarkListRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < 200; x += 5 {
			seq.Remove(x)
		}
	}
}
