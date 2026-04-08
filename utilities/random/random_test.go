// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package random

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_generate_unique_identifier_non_empty(t *testing.T) {
	id := GenerateUniqueIdentifier()
	require.NotEmpty(t, id)
}

func Test_generate_unique_identifier_two_calls_differ(t *testing.T) {
	id1 := GenerateUniqueIdentifier()
	id2 := GenerateUniqueIdentifier()
	assert.NotEqual(t, id1, id2, "two calls should produce different identifiers")
}

func Test_generate_unique_identifier_with_prefix(t *testing.T) {
	prefix := "test-prefix-"
	id := GenerateUniqueIdentifier(prefix)
	assert.True(t, strings.HasPrefix(id, prefix), "identifier should start with the given prefix")
}

func Test_generate_unique_identifier_with_multiple_prefixes(t *testing.T) {
	id := GenerateUniqueIdentifier("a", "b", "c")
	assert.True(t, strings.HasPrefix(id, "abc"), "multiple prefixes should be joined")
}

func Test_generate_unique_identifier_list(t *testing.T) {
	ids := GenerateUniqueIdentifierList("item-", 5)
	require.Len(t, ids, 5)

	seen := make(map[string]bool)
	for _, id := range ids {
		assert.True(t, strings.HasPrefix(id, "item-"))
		assert.False(t, seen[id], "identifiers in list should be unique")
		seen[id] = true
	}
}

func Test_generate_unique_identifier_list_zero_count(t *testing.T) {
	ids := GenerateUniqueIdentifierList("x-", 0)
	assert.Len(t, ids, 0)
}

func Test_generate_random_bool(t *testing.T) {
	// Just verify it does not panic and returns a bool
	_ = GenerateRandomBool()
}

func Test_generate_random_bool_distribution(t *testing.T) {
	// Run enough times to confirm both true and false appear
	trueCount := 0
	for i := 0; i < 100; i++ {
		if GenerateRandomBool() {
			trueCount++
		}
	}
	assert.Greater(t, trueCount, 0, "should produce at least one true in 100 calls")
	assert.Less(t, trueCount, 100, "should produce at least one false in 100 calls")
}
