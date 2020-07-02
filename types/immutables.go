package types

import (
	"crypto/sha1"
	"encoding/base64"
	"sort"
	"strings"
)

type Immutables interface {
	Get() Properties
	GetHashID() ID
}

type immutables struct {
	Properties Properties
}

var _ Immutables = (*immutables)(nil)

func (immutables immutables) Get() Properties {
	return immutables.Properties
}
func (immutables immutables) GetHashID() ID {
	var facts []string
	for _, immutableProperty := range immutables.Properties.GetList() {
		facts = append(facts, immutableProperty.GetFact().String())
	}
	sort.Strings(facts)
	toDigest := strings.Join(facts, "_")
	h := sha1.New()
	h.Write([]byte(toDigest))
	return NewID(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}
func NewImmutables(properties Properties) Immutables {
	return immutables{Properties: properties}
}
