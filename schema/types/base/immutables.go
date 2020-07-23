package base

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"sort"
	"strings"
)

type immutables struct {
	Properties types.Properties
}

var _ types.Immutables = (*immutables)(nil)

func (immutables immutables) Get() types.Properties {
	return immutables.Properties
}
func (immutables immutables) GetHashID() types.ID {
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
func NewImmutables(properties types.Properties) types.Immutables {
	return immutables{Properties: properties}
}
