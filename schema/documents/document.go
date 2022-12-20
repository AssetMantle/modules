// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type Document interface {
	GenerateHashID() ids.ID
	GetClassificationID() *baseIDs.ID
	// GetProperty returns property from a document searching in both Mutables and Immutables
	// * Returns nil if property is not found
	GetProperty(ids.ID) properties.Property
	GetImmutables() *baseQualified.Immutables
	GetMutables() *baseQualified.Mutables

	Mutate(...properties.Property) Document
}
