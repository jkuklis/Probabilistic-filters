package my_main

import (
	"os"
	"fmt"
	"math/rand"
	"testing"
	"github.com/linvon/cuckoo-filter"
	// "github.com/seiflotfy/cuckoofilter"
	// "github.com/panmari/cuckoofilter"
)

func BenchmarkNewCuckoo(b *testing.B) {
	var sizes = []uint {10_000, 100_000, 250_000}
	var bucket_size = []uint {4}
	var fingerprints = []uint {16}

	for _, s := range sizes {
		for _, bu := range bucket_size {
			for _, f := range fingerprints {
				b.Run(fmt.Sprintf("s_%d_b_%d_f_%d", s, bu, f), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						// cuckoo.NewFilter(bu, f, s, cuckoo.TableTypePacked)
						// cuckoo.NewFilter(bu, f, s, cuckoo.TableTypeSingle)
						// cuckoo.NewFilter(s)
					}
				})
			}
		}
	}
}

func BenchmarkFillCuckoo(b *testing.B) {
	var sizes = []uint {10_000, 100_000, 250_000}
	var bucket_size = []uint {4}
	var fingerprints = []uint {16}

	var res bool

	for _, s := range sizes {
		for _, bu := range bucket_size {
			for _, f := range fingerprints {
				// filter := cuckoo.NewFilter(bu, f, s, cuckoo.TableTypePacked)
				// filter := cuckoo.NewFilter(bu, f, s, cuckoo.TableTypeSingle)
				// filter := cuckoo.NewFilter(s)
				b.Run(fmt.Sprintf("s_%d_b_%d_f_%d", s, bu, f), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						token := make([]byte, 4)
    					rand.Read(token)
						// filter.AddUnique(token)
						// filter.Insert(token)
					}
				})
				// res = filter.Contain([]byte{0, 0, 0, 0})
			}
		}
	}
	fmt.Println(res)
}

func BenchmarkFullCreationCuckoo(b *testing.B) {
	var sizes = []uint {10_000, 100_000, 250_000}
	var bucket_size = []uint {4}
	var fingerprints = []uint {16}

	var res bool

	for _, s := range sizes {
		for _, bu := range bucket_size {
			for _, f := range fingerprints {
				b.Run(fmt.Sprintf("s_%d_b_%d_f_%d", s, bu, f), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						filter := cuckoo.NewFilter(bu, f, s, cuckoo.TableTypeSingle)
						// filter := cuckoo.NewFilter(s)
						var j uint
						for j = 0; j < s; j++ {
							token := make([]byte, 4)
							rand.Read(token)
							filter.AddUnique(token)
							// filter.Insert(token)
						}
						res = filter.Contain([]byte{0, 0, 0, 0})
						// res = filter.Lookup([]byte{0, 0, 0, 0})
					}
				})
			}
		}
	}
	fmt.Println(res)
}

func BenchmarkDeserializationCuckoo(b *testing.B) {
	var sizes = []uint {10_000, 100_000, 250_000}
	var bucket_size = []uint {4}
	var fingerprints = []uint {16}

	var res bool

	for _, s := range sizes {
		for _, bu := range bucket_size {
			for _, f := range fingerprints {
				b.Run(fmt.Sprintf("s_%d_b_%d_f_%d", s, bu, f), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						dat, err := os.ReadFile(fmt.Sprintf("tmp/cuckoo_3_size_%d_b_%d_f_%d", s, bu, f))
						Check(err)
						filter, err := cuckoo.Decode(dat)
						Check(err)
						res = filter.Contain([]byte{0, 0, 0, 0})
					}
				})
			}
		}
	}
	fmt.Println(res)
}