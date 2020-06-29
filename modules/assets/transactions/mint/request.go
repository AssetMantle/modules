package mint

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"strings"
)

type transactionRequest struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	ClassificationID string       `json:"classificationID"`
	MaintainersID    string       `json:"maintainersID"`
	Properties       string       `json:"properties"`
	Lock             int          `json:"lock"`
	Burn             int          `json:"burn"`
}

var _ types.TransactionRequest = (*transactionRequest)(nil)

func (request transactionRequest) FromCLI(cliCommand types.CLICommand, cliContext context.CLIContext) types.TransactionRequest {
	request.BaseReq = cliCommand.ReadBaseReq(cliContext)
	request.ClassificationID = cliCommand.ReadString(constants.ClassificationID)
	request.MaintainersID = cliCommand.ReadString(constants.MaintainersID)
	request.Properties = cliCommand.ReadString(constants.Properties)
	request.Lock = cliCommand.ReadInt(constants.Lock)
	request.Burn = cliCommand.ReadInt(constants.Burn)
	return request
}

func (request transactionRequest) GetBaseReq() rest.BaseReq {
	return request.BaseReq
}

func (request transactionRequest) MakeMsg() sdkTypes.Msg {
	from, Error := sdkTypes.AccAddressFromBech32(request.GetBaseReq().From)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	properties := strings.Split(request.Properties, constants.PropertiesSeparator)
	if len(properties) > constants.MaxTraitCount {
		panic(errors.New(fmt.Sprintf("")))
	}

	var propertyList []types.Property
	for _, property := range properties {
		traitIDAndProperty := strings.Split(property, constants.TraitIDAndPropertySeparator)
		if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
			propertyList = append(propertyList, types.NewProperty(types.NewID(traitIDAndProperty[0]), types.NewFact(traitIDAndProperty[1], types.NewSignatures(nil))))
		}
	}

	message := Message{
		From:             from,
		ChainID:          types.NewID(request.GetBaseReq().ChainID),
		MaintainersID:    types.NewID(request.MaintainersID),
		ClassificationID: types.NewID(request.ClassificationID),
		Properties:       types.NewProperties(propertyList),
		Lock:             types.NewHeight(request.Lock),
		Burn:             types.NewHeight(request.Burn),
	}
	return message
}

func requestPrototype() types.TransactionRequest {
	return &transactionRequest{}
}
