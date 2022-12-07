// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ documents.Document = (*Document)(nil)

func (document *Document) GenerateHashID() ids.HashID {
	return base.GenerateHashID(document.GetClassificationID().Bytes(), document.GetImmutables().GenerateHashID().Bytes())
}
func (document *Document) GetClassificationID() ids.ClassificationID {
	return document.ClassificationId
}
func (document *Document) GetProperty(propertyID ids.PropertyID) properties.Property {
	if property := document.Immutables.GetImmutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else if property := document.Document.Mutables.GetMutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else {
		return nil
	}
}
func (document *Document) GetImmutables() qualified.Immutables {
	return document.Immutables
}
func (document *Document) GetMutables() qualified.Mutables {
	return document.Mutables
}

// TODO write test case
func (document *Document) Mutate(propertyList ...properties.Property) documents.Document {
	//document.Document.Mutables = document.Document.Mutables.Mutate(propertyList...)
	return document
}

func NewDocument(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Document {
	return &Document{
		ClassificationId: classificationID,
		Immutables:       immutables,
		Mutables:         mutables,
	}
}
