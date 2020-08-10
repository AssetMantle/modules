/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

//TODO add mutable flag to traits
type transactionRequest struct {
	BaseReq       rest.BaseReq `json:"baseReq"`
	FromID        string       `json:"fromID" valid:"required~required field fromID missing"`
	MaintainersID string       `json:"maintainersID" valid:"required~required field maintainersID missing matches(^[A-Za-z]$)~invalid field maintainersID"`
	Traits        string       `json:"traits" valid:"required~required field traits missing matches(^[A-Za-z]$)~invalid field traits"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadString(constants.MaintainersID),
		cliCommand.ReadString(constants.Traits),
	)
}

func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}

func (transactionRequest transactionRequest) MakeMsg() sdkTypes.Msg {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	traits := strings.Split(transactionRequest.Traits, constants.TraitsSeparator)
	if len(traits) > constants.MaxTraitCount {
		//TODO handle
		panic(errors.New(fmt.Sprintf("")))
	}

	var traitList []types.Trait
	for _, trait := range traits {
		traitIDAndProperty := strings.Split(trait, constants.PropertyIDAndFactSeparator)
		if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
			traitID := base.NewID(traitIDAndProperty[0])
			//TODO check proper working
			traitList = append(traitList, base.NewTrait(base.NewProperty(traitID, base.NewFact(traitIDAndProperty[1])), true))
		}
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.MaintainersID),
		base.NewTraits(traitList),
	)
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, maintainersID string, traits string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:       baseReq,
		FromID:        fromID,
		MaintainersID: maintainersID,
		Traits:        traits,
	}
}
