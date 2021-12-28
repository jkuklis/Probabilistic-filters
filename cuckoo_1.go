package main

import (
	"os"
	"fmt"
	"github.com/seiflotfy/cuckoofilter"
)

func CreateCuckoo1() {
	var sizes = []struct {
		size uint
	}{
		{size: 10_000},
		{size: 100_000},
		{size: 250_000},
	}

	for _, s := range sizes {
		filter := cuckoo.NewFilter(s.size)
		filename := fmt.Sprintf("tmp/cuckoo_1_size_%d_fp_0.03", s.size)
		err := os.WriteFile(filename, filter.Encode(), 0644)
		Check(err)
	}
}
