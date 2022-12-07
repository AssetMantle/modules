// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
	base2 "github.com/AssetMantle/modules/schema/qualified/base"
)

var _ documents.Document = (*DocumentI_Document)(nil)

func (document *DocumentI_Document) GenerateHashID() ids.HashID {
	return base.GenerateHashID(document.GetClassificationID().Bytes(), document.GetImmutables().GenerateHashID().Bytes())
}
func (document *DocumentI_Document) GetClassificationID() ids.ClassificationID {
	return document.Document.ClassificationId
}
func (document *DocumentI_Document) GetProperty(propertyID ids.PropertyID) properties.Property {
	if property := document.Document.Immutables.GetImmutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else if property := document.Document.Mutables.GetMutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else {
		return nil
	}
}
func (document *DocumentI_Document) GetImmutables() qualified.Immutables {
	return document.Document.Immutables
}
func (document *DocumentI_Document) GetMutables() qualified.Mutables {
	return document.Document.Mutables
}

// TODO write test case
func (document *DocumentI_Document) Mutate(propertyList ...properties.Property) documents.Document {
	document.Document.Mutables = document.Document.Mutables.Mutate(propertyList...)
	return document
}

func NewDocument(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Document {
	return &DocumentI{
		Impl: &DocumentI_Document{
			Document: &Document{
				ClassificationId: classificationID.(*base.ClassificationIDI),
				Immutables:       immutables.(*base2.Immutables),
				Mutables:         mutables.(*base2.Mutables),
			},
		},
	}
}
