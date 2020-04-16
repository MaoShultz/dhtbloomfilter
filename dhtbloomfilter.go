package dhtbloomfilter

import (
	"encoding/hex"
	"math"
	"math/bits"
)

const bfSize = 256 // in bytes

/*
BloomFilter sha1 base use in DHT get_peers response. BEP33
*/
type BloomFilter [bfSize]byte

/*
Merge сombines two filters
*/
func (bf *BloomFilter) Merge(filter BloomFilter) {
	for i := 0; i < bfSize; i++ {
		bf[i] = bf[i] | filter[i]
	}
}

/*
EstimatedSize Rounded integer, approximating the number of items in the filter.
*/
func (bf *BloomFilter) EstimatedSize() int {

	zeros := 0

	for i := 0; i < bfSize; i++ {
		zeros = zeros + bits.OnesCount8(^bf[i])
	}

	if zeros == 0 {
		return 6000 // The maximum capacity of the bloom filter used in BEP33
	}

	m := 256 * 8
	c := math.Min(float64(m-1), float64(zeros))

	return int(math.Log(c/float64(m)) / (2 * math.Log(1-1/float64(m))))
}

/*
Dump bytes as HEX string
*/
func (bf *BloomFilter) Dump() string {
	return hex.Dump(bf)
}
