package my_main

import (
	"os"
	"fmt"
	"github.com/bits-and-blooms/bloom"
)

func CreateBloom1() {
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
			filter := bloom.NewWithEstimates(s.size, fp.fp)
			filename := fmt.Sprintf("tmp/bloom_1_size_%d_fp_%.3f", s.size, fp.fp)
			f, err := os.Create(filename)
			Check(err)
			_, err = filter.WriteTo(f)
			Check(err)
		}
	}
}
