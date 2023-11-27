package gobloom

import (
	"fmt"
	"hash"
	"math"
	"sync"
)

type BloomFilter struct {
	bitSet        []bool
	bitSetLength  uint64        //short for len(bitSet)
	hashes        []hash.Hash64 //hash functions to use
	numOfHashFunc uint64        //too few hash func lead to crowded data/ false pos, too many slows down filter/performance
	mutex         sync.Mutex    //ensure thread safety
}

func NewBloomFilterWithHasher(numOfElemToStore uint64,
	acceptableFalsePosRate float64,
	hasher Hasher) (*BloomFilter, error) {

	if numOfElemToStore == 0 {
		return nil, fmt.Errorf("must store at least one element")
	}
	if acceptableFalsePosRate <= 0 || acceptableFalsePosRate >= 1 {
		return nil, fmt.Errorf("false positive rate must be between 0 and 1")
	}
	if hasher == nil {
		return nil, fmt.Errorf("hasher cannot be nil")
	}

	bitSetLength, numOfHashFunc := getOptimalParams(numOfElemToStore, acceptableFalsePosRate)
	return &BloomFilter{
		bitSetLength:  bitSetLength,
		numOfHashFunc: numOfHashFunc,
		bitSet:        make([]bool, bitSetLength),
		hashes:        hasher.GetHashes(numOfHashFunc),
	}, nil
}

func getOptimalParams(numOfElemToStore uint64, falsePosRate float64) (uint64, uint64) {
	m := uint64(math.Ceil(-1 * float64(numOfElemToStore) * math.Log(falsePosRate) / math.Pow(math.Log(2), 2)))
	if m == 0 {
		m = 1
	}
	k := uint64(math.Ceil((float64(m) / float64(numOfElemToStore)) * math.Log(2)))
	if k == 0 {
		k = 1
	}

	return m, k
}

func (bf *BloomFilter) Add(data []byte) {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()
	for _, hash := range bf.hashes {
		hash.Reset()
		hash.Write(data)
		hashValue := hash.Sum64() % bf.bitSetLength
		bf.bitSet[hashValue] = true
	}
}

func (bf *BloomFilter) Test(data []byte) bool {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()
	for _, hash := range bf.hashes {
		hash.Reset()
		hash.Write(data)
		hashValue := hash.Sum64() % bf.bitSetLength
		if !bf.bitSet[hashValue] {
			return false
		}
	}
	return true
}
