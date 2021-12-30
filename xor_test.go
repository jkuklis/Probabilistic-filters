package my_main

import (
	"os"
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

func TestBinFuse8(t *testing.T) {
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

	for i, _ := range sizes {
		filter, err := xorfilter.PopulateBinaryFuse8(keys[i])
		Check(err)
		fmt.Printf("Seed: %d\n", filter.Seed)
		fmt.Printf("SegmentLength: %d\n", filter.SegmentLength)
		fmt.Printf("SegmentLengthMask: %d\n", filter.SegmentLengthMask)
		fmt.Printf("SegmentCount: %d\n", filter.SegmentCount)
		fmt.Printf("SegmentCountLength: %d\n", filter.SegmentCountLength)
		fmt.Printf("\n")
	}
}

func BenchmarkDeserializationXor(b *testing.B) {
	var sizes = []uint {10_000, 100_000, 250_000}
	var res bool

	var seed = uint64(10451216379200822465)
	var segmentLength = []uint32 {512, 2048, 4096}
	var segmentLengthMask = []uint32 {511, 2047, 4095}
	var segmentCount = []uint32 {23, 56, 69}
	var segmentCountLength = []uint32 {11776, 114688, 282624}

	for i, s := range sizes {
		b.Run(fmt.Sprintf("s_%d", s), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				dat, err := os.ReadFile(fmt.Sprintf("tmp/xor_size_%d_fp_0.003", s))
				Check(err)
				var filter = xorfilter.BinaryFuse8 {
					seed,
					segmentLength[i],
					segmentLengthMask[i],
					segmentCount[i],
					segmentCountLength[i],
					dat,
				}
				Check(err)
				res = filter.Contains(0)
			}
		})
	}
	fmt.Println(res)
}