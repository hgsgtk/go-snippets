package bloomfilter

import (
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	m uint64 // number of bits in the filter
	k uint8  // number of hash functions
	bits []byte // bit array
}

func NewBloomFilter(size int, errorRate float64) *BloomFilter {
	m, k := estimateParameters(size, errorRate)

	return &BloomFilter{
		m: m,
		k: k,
		bits: make([]byte, (m+7)/8),
	}
}

func estimateParameters(size int, errorRate float64) (uint64, uint8) {
	if size == 0 || errorRate <= 0 || errorRate >= 1 {
		return 1000, 10
	}

	m := uint64(math.Ceil(-1 * float64(size) * math.Log2(errorRate)))
	k := uint8(math.Ceil(math.Log2(errorRate)))

	return m, k
}

func (bf *BloomFilter) hash(data []byte) []uint64 {
	hashes := make([]uint64, bf.k)

	h1 := fnv.New64a()
	h1.Write(data)
	hash1Val := h1.Sum64()

	h2 := fnv.New64a()
	h2.Write(data)
	hash2Val := h2.Sum64()

	for i := uint8(0); i < bf.k; i++ {
		if hash2Val == 0 && i > 0 {
			hash2Val = 1
		}
		hashes[i] = (hash1Val + uint64(i)*hash2Val) % bf.m
	}
	return hashes
}

func (bf *BloomFilter) Add(item string) {
	hashes := bf.hash([]byte(item))

	for _, hash := range hashes {
		byteIndex := hash / 8
		bitOffset := hash % 8
		bf.bits[byteIndex] |= 1 << bitOffset
	}
}

func (bf *BloomFilter) Contains(item string) bool {
	hashes := bf.hash([]byte(item))

	for _, hash := range hashes {
		byteIndex := hash / 8
		bitOffset := hash % 8
		if (bf.bits[byteIndex] & (1 << bitOffset)) == 0 {
			return false
		}
	}
	return true
}
