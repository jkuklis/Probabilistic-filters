package my_main

import (
	"os"
	"fmt"
	"github.com/linvon/cuckoo-filter"
)

func CreateCuckoo3() {
	var sizes = []struct {
		size uint
	}{
		{size: 10_000},
		{size: 100_000},
		{size: 250_000},
	}

	var bytes = []uint {4}
	var fingerprints = []uint {8, 9, 10, 16}

	for _, s := range sizes {
		for _, b := range bytes {
			for _, f := range fingerprints {
				filter := cuckoo.NewFilter(b, f, s.size, cuckoo.TableTypePacked)
				filename := fmt.Sprintf("tmp/cuckoo_3_size_%d_b_%d_f_%d", s.size, b, f)
				bytes, err := filter.Encode()
				Check(err)
				err = os.WriteFile(filename, bytes, 0644)
				Check(err)
			}
		}
	}
}
