package cuckoo

import (
	"crypto/sha1"
	"encoding/binary"
	"math"
)

type bucket []fingerprint
type fingerprint []byte

var hasher = sha1.New()

type Cuckoo struct {
	buckets []bucket
	m       uint // buckets
	b       uint // entries per bucket
	f       uint // fingerprint length
	n       uint // filter capacity (rename cap?)
}

func NewCuckoo(n uint, fp float64) *Cuckoo {
	b := uint(4)
	f := fingerprintLength(b, fp)
	m := nextPower(n / f * 8)
	buckets := make([]bucket, m)
	for i := uint(0); i < m; i++ {
		buckets[i] = make(bucket, b)
	}
	return &Cuckoo{
		buckets: buckets,
		m:       m,
		b:       b,
		f:       f,
		n:       n,
	}
}


func (c *Cuckoo) hashes(data string) (uint, uint, fingerprint) {
	h := hash([]byte(data))
	f := h[0:c.f]
	i1 := uint(binary.BigEndian.Uint32(h))
	i2 := i1 ^ uint(binary.BigEndian.Uint32(hash(f)))
	return i1, i2, fingerprint(f)
}


func hash(data []byte) []byte {
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)
	hasher.Reset()
	return hash
}


func fingerprintLength(b uint, e float64) uint {
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
