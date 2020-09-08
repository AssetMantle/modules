/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type classifications struct {
	ID   types.ID                   `json:"id" valid:"required~required field id missing"`
	List []mappables.Classification `json:"list" valid:"required~required list missing"`

	mapper  helpers.Mapper
	context sdkTypes.Context
}

var _ mappers.Classifications = (*classifications)(nil)

func (classifications classifications) GetID() types.ID { return classifications.ID }
func (classifications classifications) Get(id types.ID) mappables.Classification {
	classificationID := classificationIDFromInterface(id)
	for _, oldClassification := range classifications.List {
		if oldClassification.GetID().Equal(classificationID) {
			return oldClassification
		}
	}
	return nil
}
func (classifications classifications) GetList() []mappables.Classification {
	return classifications.List
}

func (classifications classifications) Fetch(id types.ID) mappers.Classifications {
	var classificationList []mappables.Classification
	classificationsID := classificationIDFromInterface(id)
	if len(classificationsID.HashID.Bytes()) > 0 {
		mappable := classifications.mapper.Read(classifications.context, classificationsID)
		if mappable != nil {
			classificationList = append(classificationList, mappable.(classification))
		}
	} else {
		appendClassificationList := func(mappable traits.Mappable) bool {
			classificationList = append(classificationList, mappable.(classification))
			return false
		}
		classifications.mapper.Iterate(classifications.context, classificationsID, appendClassificationList)
	}
	classifications.ID, classifications.List = id, classificationList
	return classifications
}
func (classifications classifications) Add(classification mappables.Classification) mappers.Classifications {
	classifications.ID = readClassificationID("")
	classifications.mapper.Create(classifications.context, classification)
	classifications.List = append(classifications.List, classification)
	return classifications
}
func (classifications classifications) Remove(classification mappables.Classification) mappers.Classifications {
	classifications.mapper.Delete(classifications.context, classification.GetID())
	for i, oldClassification := range classifications.List {
		if oldClassification.GetID().Equal(classification.GetID()) {
			classifications.List = append(classifications.List[:i], classifications.List[i+1:]...)
			break
		}
	}
	return classifications
}
func (classifications classifications) Mutate(classification mappables.Classification) mappers.Classifications {
	classifications.mapper.Update(classifications.context, classification)
	for i, oldClassification := range classifications.List {
		if oldClassification.GetID().Equal(classification.GetID()) {
			classifications.List[i] = classification
			break
		}
	}
	return classifications
}

func NewClassifications(mapper helpers.Mapper, context sdkTypes.Context) mappers.Classifications {
	return classifications{
		ID:      readClassificationID(""),
		List:    []mappables.Classification{},
		mapper:  mapper,
		context: context,
	}
}
