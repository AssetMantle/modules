/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package conform

import (
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	fmt.Println("Classifications calc")
	fmt.Println(key.FromID(auxiliaryRequest.ClassificationID), "Printing classificationID")
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.FromID(auxiliaryRequest.ClassificationID))

	fmt.Println("Done classifications Calc")
	classification := classifications.Get(key.FromID(auxiliaryRequest.ClassificationID))
	fmt.Println("singular calc", classification)
	if classification == nil {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}

	fmt.Println(auxiliaryRequest.ImmutableProperties, "Printing Immutable Properties")
	if auxiliaryRequest.ImmutableProperties != nil {
		fmt.Println("If imm != nil")
		if len(auxiliaryRequest.ImmutableProperties.GetList()) != len(classification.(mappables.Classification).GetImmutableProperties().GetList()) {
			fmt.Println("%%%%%%%")
			return newAuxiliaryResponse(errors.NotAuthorized)
		}

		fmt.Println("Entering For loop")
		for _, immutableProperty := range auxiliaryRequest.ImmutableProperties.GetList() {
			if property := classification.(mappables.Classification).GetImmutableProperties().Get(immutableProperty.GetID()); property == nil || property.GetFact().GetTypeID().Compare(immutableProperty.GetFact().GetTypeID()) != 0 || property.GetFact().GetHashID().Compare(base.NewID("")) != 0 && property.GetFact().GetHashID() != immutableProperty.GetFact().GetHashID() {
				fmt.Println("Going thorugh for loop")
				return newAuxiliaryResponse(errors.NotAuthorized)
			}
		}
	}

	fmt.Println("Done with prev if")
	if auxiliaryRequest.MutableProperties != nil {
		fmt.Println(len(auxiliaryRequest.MutableProperties.GetList()), "Mutable Properties Print")
		fmt.Println(len(classification.(mappables.Classification).GetMutableProperties().GetList()), "Second")
		if len(auxiliaryRequest.MutableProperties.GetList()) > len(classification.(mappables.Classification).GetMutableProperties().GetList()) {
			fmt.Println("%%%%%%%")
			return newAuxiliaryResponse(errors.NotAuthorized)
		}

		fmt.Println("Entering For loop")
		for _, mutableProperty := range auxiliaryRequest.MutableProperties.GetList() {
			if property := classification.(mappables.Classification).GetMutableProperties().Get(mutableProperty.GetID()); property == nil || property.GetFact().GetTypeID().Compare(mutableProperty.GetFact().GetTypeID()) != 0 {
				fmt.Println("Going thorugh for lopp")
				return newAuxiliaryResponse(errors.NotAuthorized)
			}
		}
	}

	fmt.Println("Finish")
	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
