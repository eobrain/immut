package slice_test

import (
	"github.com/eobrain/immut/slice"
	"math/rand"
	"testing"
)

var seq = slice.New()

/////////////////////////////////////////////////////////////////////////////

func init() {
	for i := 0; i < 100; i++ {
		seq = seq.AddFront(rand.Int())
	}
}

func BenchmarkRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < 200; x += 5 {
			seq.Remove(x)
		}
	}
}

func BenchmarkFront(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq.Front()
	}
}

func BenchmarkBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq.Back()
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq.Get(50)
	}
}

func BenchmarkAddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq.AddAll(seq)
	}
}

func BenchmarkAddAll_huge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		huge := seq
		for i := 0; i < 20; i++ {
			huge.AddAll(huge)
		}
	}
}
