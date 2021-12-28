package main

import (
	"os"
	"fmt"
	"github.com/FastFilter/xorfilter"
)

func CreateXor() {
	var sizes = []struct {
		size uint64
	}{
		{size: 10_000},
		{size: 100_000},
		{size: 250_000},
	}

	var keys [][]uint64

	for _, s := range sizes {
		data := make([]uint64, s.size)
		var i uint64
		for i = 0; i < s.size; i++ {
			data[i] = i
		}
		keys = append(keys, data)
	}

	for i, s := range sizes {
		filter, err := xorfilter.PopulateBinaryFuse8(keys[i])
		Check(err)
		filename := fmt.Sprintf("tmp/xor_size_%d_fp_0.003", s.size)
		// should also serialize other attributes, 24 bytes
		err = os.WriteFile(filename, filter.Fingerprints, 0644)
		Check(err)
	}
}
