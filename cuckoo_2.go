package main

import (
	"os"
	"fmt"
	"github.com/panmari/cuckoofilter"
)

func CreateCuckoo2() {
	var sizes = []struct {
		size uint
	}{
		{size: 10_000},
		{size: 100_000},
		{size: 250_000},
	}

	for _, s := range sizes {
		filter := cuckoo.NewFilter(s.size)
		filename := fmt.Sprintf("tmp/cuckoo_2_size_%d_fp_0.0001", s.size)
		err := os.WriteFile(filename, filter.Encode(), 0644)
		Check(err)
	}
}
