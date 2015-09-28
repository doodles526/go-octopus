package valhash

import (
	"bytes"
	"crypto/sha256"
)

type ValHash struct {
	Value  []byte
	hashID []byte
}

func NewValHash(val []byte) *ValHash {
	hash := sha256.Sum256(val)
	return &ValHash{Value: val, hashID: hash[:sha256.Size]}
}

func (v *ValHash) Compare(v2 *ValHash) int {
	return bytes.Compare(v.hashID, v2.hashID)
}

// NewKeyValWithHash allows us to set the hash
// Value. This is only to be used in tests
func NewValHashWithHash(val []byte, hash []byte) *ValHash {
	return &ValHash{Value: val, hashID: hash}
}
