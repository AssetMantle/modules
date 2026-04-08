// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/queries/identities"
	"github.com/AssetMantle/modules/x/identities/queries/identity"
	"github.com/AssetMantle/modules/x/identities/queries/parameters"
	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Queries
	}{
		{"valid", baseHelpers.NewQueries(
			identity.Query,
			identities.Query,
			parameters.Query,
		)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prototype()
			assert.NotNil(t, got)
		})
	}
}
