package gobloom

import (
	"hash"
)

type Interface interface { //foundational function of any bloom filter
	// add an element then
	//test if it is there

	Add([]byte)
	Test([]byte) bool
}

type Hasher interface { //make hashes for bloom filter impl
	//separate hash interface allows experimentation with diff
	//hash functions to balance collision, comp efficiency, etc.
	GetHashes(n uint64) []hash.Hash64
}
