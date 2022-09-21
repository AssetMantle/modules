// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

type Document struct {
	ID ids.ID `json:"id" valid:"required~required field id is missing"`
	Immutables
	Mutables //nolint:govet
}

var _ qualified.Document = (*Document)(nil)

func (document Document) GetID() ids.ID {
	return document.ID
}
func (document Document) GetClassificationID() ids.ID {
	panic("no implemented ")
}
func (document Document) GetProperty(propertyID ids.PropertyID) properties.Property {
	if property := document.Immutables.GetImmutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else if property := document.Mutables.GetMutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else {
		return nil
	}
}

// TODO write test case
func (document Document) Mutate(propertyList ...properties.Property) qualified.Document {
	document.Mutables = document.Mutables.Mutate(propertyList...).(Mutables)
	return document
}
