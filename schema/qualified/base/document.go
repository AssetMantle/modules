// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
)

type Document struct {
	ID               types.ID `json:"id" valid:"required~required field id is missing"`
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID is missing"`
	Immutables
	Mutables //nolint:govet
}

var _ qualified.Document = (*Document)(nil)

func (document Document) GetID() types.ID {
	return document.ID
}
func (document Document) GetClassificationID() types.ID {
	return document.ClassificationID
}
func (document Document) GetProperty(propertyID ids.PropertyID) types.Property {
	if property := document.Immutables.GetImmutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else if property := document.Mutables.GetMutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else {
		return nil
	}
}

// TODO write test case
func (document Document) Mutate(propertyList ...types.Property) qualified.Document {
	document.Mutables = document.Mutables.Mutate(propertyList...).(Mutables)
	return document
}
