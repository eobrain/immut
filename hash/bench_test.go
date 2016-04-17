package hash_test

import (
	"github.com/eobrain/immut/hash"
	"math/rand"
	"testing"
)

var seq = hash.New()

/////////////////////////////////////////////////////////////////////////////

func init() {
	for i := 0; i < 1000; i++ {
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

func BenchmarkRest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq.Rest()
	}
}

func BenchmarkBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq.Back()
	}
}

func BenchmarkLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq.Len()
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
