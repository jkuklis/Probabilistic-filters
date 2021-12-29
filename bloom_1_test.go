package main

import (
	"fmt"
	"testing"
	"github.com/bits-and-blooms/bloom"
)

func BenchmarkCreateBloom1(b *testing.B) {
	var sizes = []struct {
		size uint
	}{
		{size: 10_000},
		{size: 100_000},
		{size: 250_000},
	}

	var fp_rates = []struct {
		fp float64
	}{
		{fp: 0.01},
		{fp: 0.001},
	}

	for _, s := range sizes {
		for _, fp := range fp_rates {
			b.Run(fmt.Sprintf("size_%d_fp_%.2f", s.size, fp.fp), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					bloom.NewWithEstimates(s.size, fp.fp)
				}
			})
		}
	}
}
