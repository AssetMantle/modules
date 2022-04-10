// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package cuckoo

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"math"
	"math/rand"
)

type bucket []fingerprint
type fingerprint []byte

// Cuckoo https://www.pdl.cmu.edu/PDL-FTP/FS/cuckoo-conext2014.pdf
type Cuckoo struct {
	buckets           []bucket
	numBuckets        uint // buckets
	entriesPerBucket  uint // entries per bucket
	fingerprintLength uint // fingerprint length
	n                 uint // filter capacity (rename cap?)
	retries           int  // how many times do we try to move items around?
}

func NewCuckoo(numItems uint, entriesPerBucket uint, retries int, falsePositiveRate float64) *Cuckoo {
	fingerprintLength := getFingerprintLength(entriesPerBucket, falsePositiveRate)
	numBuckets := nextPower(numItems / fingerprintLength * 8)
	buckets := make([]bucket, numBuckets)

	for i := uint(0); i < numBuckets; i++ {
		buckets[i] = make(bucket, entriesPerBucket)
	}

	return &Cuckoo{
		buckets:           buckets,
		numBuckets:        numBuckets,
		entriesPerBucket:  entriesPerBucket,
		fingerprintLength: fingerprintLength,
		n:                 numItems,
		retries:           retries,
	}
}

// Delete delete the fingerprint from the cuckoo filter
func (c *Cuckoo) Delete(needle string) {
	i1, i2, f := c.hashes(needle)
	// try to remove from f1
	b1 := c.buckets[i1%c.numBuckets]
	if ind, ok := b1.Contains(f); ok {
		b1[ind] = nil
		return
	}

	b2 := c.buckets[i2%c.numBuckets]
	if ind, ok := b2.Contains(f); ok {
		b2[ind] = nil
		return
	}
}

// Lookup find needle in the cuckoo filter
func (c *Cuckoo) Lookup(needle string) bool {
	i1, i2, f := c.hashes(needle)
	_, b1 := c.buckets[i1%c.numBuckets].Contains(f)
	_, b2 := c.buckets[i2%c.numBuckets].Contains(f)

	return b1 || b2
}

func (b bucket) Contains(f fingerprint) (int, bool) {
	for i, x := range b {
		if bytes.Equal(x, f) {
			return i, true
		}
	}

	return -1, false
}

func (c *Cuckoo) Insert(input string) error {
	i1, i2, f := c.hashes(input)
	// bucket one
	b1 := c.buckets[i1%c.numBuckets]
	if i, err := b1.nextIndex(); err == nil {
		b1[i] = f
		return nil
	}

	// bucket two
	b2 := c.buckets[i2%c.numBuckets]
	if i, err := b2.nextIndex(); err == nil {
		b2[i] = f
		return nil
	}

	// else we need to start relocating items
	i := i1
	for r := 0; r < c.retries; r++ {
		index := i % c.numBuckets
		entryIndex := rand.Intn(int(c.entriesPerBucket)) //nolint:gosec
		// swap
		f, c.buckets[index][entryIndex] = c.buckets[index][entryIndex], f

		i ^= uint(binary.BigEndian.Uint32(Hash(f)))
		b := c.buckets[i%c.numBuckets]

		if idx, err := b.nextIndex(); err == nil {
			b[idx] = f

			return nil
		}
	}

	return errors.New("bucket full")
}

// nextIndex returns the next index for entry, or an error if the bucket is full
func (b bucket) nextIndex() (int, error) {
	for i, f := range b {
		if f == nil {
			return i, nil
		}
	}

	return -1, errors.New("bucket full")
}

// hashes returns h1, h2 and the fingerprint
func (c *Cuckoo) hashes(data string) (uint, uint, fingerprint) {
	h := Hash([]byte(data))
	f := h[0:c.fingerprintLength]
	i1 := uint(binary.BigEndian.Uint32(h))
	i2 := i1 ^ uint(binary.BigEndian.Uint32(Hash(f)))

	return i1, i2, f
}

func Hash(data []byte) []byte {
	shaHash := sha256.New()
	shaHash.Write(data)
	hash := shaHash.Sum(nil)

	return hash
}

func getFingerprintLength(b uint, e float64) uint {
	f := uint(math.Ceil(math.Log(2 * float64(b) / e)))
	f /= 8

	if f < 1 {
		return 1
	}

	return f
}

func nextPower(i uint) uint {
	i--
	i |= i >> 1
	i |= i >> 2
	i |= i >> 4
	i |= i >> 8
	i |= i >> 16
	i |= i >> 32
	i++

	return i
}
