package my_main

import (
	"fmt"
	"testing"
	"github.com/FastFilter/xorfilter"
)

func BenchmarkNewBinFuse8(b *testing.B) {
	var sizes = []uint64 {10_000, 100_000, 250_000}
	var keys [][]uint64

	for _, s := range sizes {
		data := make([]uint64, s)
		var i uint64
		for i = 0; i < s; i++ {
			data[i] = i
		}
		keys = append(keys, data)
	}

	b.ResetTimer()

	for i, s := range sizes {
		b.Run(fmt.Sprintf("size_%d", s), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				xorfilter.PopulateBinaryFuse8(keys[i])
			}
		})
	}
}
