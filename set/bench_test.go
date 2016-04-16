package set_test

import (
	"github.com/eobrain/immut"
	"github.com/eobrain/immut/set"
	"testing"
)

var seq = set.New()

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
