// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/qualified/base"
)

//type document struct {
//	ids.ClassificationID
//	qualified.ImmutablesList
//	qualified.Mutables
//}

var _ documents.Document = (*Document)(nil)

func (document *Document) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(document.GetClassificationID().Bytes(), document.GetImmutables().GenerateHashID().Bytes())
}
func (document *Document) GetProperty(propertyID ids.PropertyID) properties.Property {
	if property := document.ImmutablesList.GetProperty(propertyID); property != nil {
		return property
	} else if property := document.MutablesList.GetProperty(propertyID); property != nil {
		return property
	} else {
		return nil
	}
}
func (document *Document) GetClassificationID() ids.ClassificationID {
	return document.ClassificationId
}
func (document *Document) GetImmutables() qualified.Immutables {
	return document.ImmutablesList
}
func (document *Document) GetMutables() qualified.Mutables {
	return document.MutablesList
}

// TODO write test case
func (document *Document) Mutate(propertyList ...properties.Property) documents.Document {
	document.MutablesList = document.MutablesList.Mutate(propertyList...).(*base.Mutables)
	return document
}

func NewDocument(classificationID ids.ID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Document {
	return &Document{
		ClassificationId: classificationID.(*baseIDs.AnyID),
		ImmutablesList:   immutables.(*base.Immutables),
		MutablesList:     mutables.(*base.Mutables),
	}
}
