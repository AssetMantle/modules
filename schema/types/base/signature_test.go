// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/base64"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

type testableStringID struct {
	IDString string
}

func (t testableStringID) AsString() string {
	// TODO implement me
	panic("implement me")
}

func (t testableStringID) ToAnyID() ids.AnyID {
	// TODO implement me
	panic("implement me")
}

func (t testableStringID) Compare(listable traits.Listable) int {
	// TODO implement me
	panic("implement me")
}

func (t testableStringID) String() string {
	// TODO implement me
	panic("implement me")
}

func (t testableStringID) Bytes() []byte {
	return []byte(t.IDString)
}

func (t testableStringID) IsStringID() {
	// TODO implement me
	panic("implement me")
}

func NewStringID(idString string) ids.StringID {
	return testableStringID{IDString: idString}
}

func TestNewSignature(t *testing.T) {

	privateKey := ed25519.GenPrivKey()
	// pubKey := privateKey.PubKey()
	signatureBytes := []byte("Temp")

	signedBytes, err := privateKey.Sign(signatureBytes)
	require.Nil(t, err)

	id := NewStringID("ID")
	validityHeight := NewHeight(123)
	testSignature := NewSignature(id, signedBytes, validityHeight)

	type args struct {
		id             ids.ID
		signatureBytes []byte
		validityHeight types.Height
	}
	tests := []struct {
		name string
		args args
		want types.Signature
	}{
		{"Test for New Signature", args{id, signedBytes, validityHeight}, testSignature},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSignature(tt.args.id, tt.args.signatureBytes, tt.args.validityHeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signature_Bytes(t *testing.T) {

	privateKey := ed25519.GenPrivKey()
	// pubKey := privateKey.PubKey()
	signatureBytes := NewStringID("Temp").Bytes()

	signedBytes, err := privateKey.Sign(signatureBytes)
	require.Nil(t, err)

	id := NewStringID("ID")
	validityHeight := NewHeight(123)
	baseSignature := NewSignature(id, signedBytes, validityHeight)

	tests := []struct {
		name          string
		baseSignature types.Signature
		want          []byte
	}{
		{"Test for Bytes", baseSignature, signedBytes},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseSignature.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signature_GetID(t *testing.T) {

	privateKey := ed25519.GenPrivKey()
	// pubKey := privateKey.PubKey()
	signatureBytes := NewStringID("Temp").Bytes()

	signedBytes, err := privateKey.Sign(signatureBytes)
	require.Nil(t, err)

	id := NewStringID("ID")
	validityHeight := NewHeight(123)
	baseSignature := NewSignature(id, signedBytes, validityHeight)

	tests := []struct {
		name          string
		baseSignature types.Signature
		want          ids.ID
	}{
		{"Test to get ID", baseSignature, id},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseSignature.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signature_GetValidityHeight(t *testing.T) {

	privateKey := ed25519.GenPrivKey()
	// pubKey := privateKey.PubKey()
	signatureBytes := NewStringID("Temp").Bytes()

	signedBytes, err := privateKey.Sign(signatureBytes)
	require.Nil(t, err)

	id := NewStringID("ID")
	validityHeight := NewHeight(123)
	baseSignature := NewSignature(id, signedBytes, validityHeight)

	tests := []struct {
		name          string
		baseSignature types.Signature
		want          types.Height
	}{
		{"Test for GetValidityHeight", baseSignature, validityHeight},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseSignature.GetValidityHeight(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValidityHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signature_HasExpired(t *testing.T) {

	privateKey := ed25519.GenPrivKey()
	// pubKey := privateKey.PubKey()
	signatureBytes := NewStringID("Temp").Bytes()

	signedBytes, err := privateKey.Sign(signatureBytes)
	require.Nil(t, err)

	id := NewStringID("ID")
	validityHeight := NewHeight(123)
	baseSignature := NewSignature(id, signedBytes, validityHeight)

	type args struct {
		height types.Height
	}
	tests := []struct {
		name          string
		baseSignature types.Signature
		args          args
		want          bool
	}{
		{"Test for Signature Expired", baseSignature, args{NewHeight(-10)}, true},
		{"Test for Signature Not Expired", baseSignature, args{validityHeight}, false},
		{"Test for zero case", baseSignature, args{NewHeight(0)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseSignature.HasExpired(tt.args.height); got != tt.want {
				t.Errorf("HasExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signature_String(t *testing.T) {

	privateKey := ed25519.GenPrivKey()
	// pubKey := privateKey.PubKey()
	signatureBytes := NewStringID("Temp").Bytes()

	signedBytes, err := privateKey.Sign(signatureBytes)
	require.Nil(t, err)

	id := NewStringID("ID")
	validityHeight := NewHeight(123)
	baseSignature := NewSignature(id, signedBytes, validityHeight)

	tests := []struct {
		name          string
		baseSignature types.Signature
		want          string
	}{
		{"Test for String", baseSignature, base64.URLEncoding.EncodeToString(baseSignature.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseSignature.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signature_Verify(t *testing.T) {

	privateKey := ed25519.GenPrivKey()
	pubKey := privateKey.PubKey()
	signatureBytes := NewStringID("Temp").Bytes()

	signedBytes, err := privateKey.Sign(signatureBytes)
	require.Nil(t, err)

	id := NewStringID("ID")
	validityHeight := NewHeight(123)
	baseSignature := NewSignature(id, signedBytes, validityHeight)

	type args struct {
		pubKey crypto.PubKey
		bytes  []byte
	}
	tests := []struct {
		name          string
		baseSignature types.Signature
		args          args
		want          bool
	}{
		{"Test for Verify Signature", baseSignature, args{pubKey, signatureBytes}, true},
		{"Test for Not Equal case", baseSignature, args{pubKey, []byte{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseSignature.Verify(tt.args.pubKey, tt.args.bytes); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
