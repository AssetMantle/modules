package mint

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type request struct {
	BaseReq          rest.BaseReq         `json:"baseReq"`
	ClassificationID types.BaseID         `json:"classificationID"`
	MaintainersID    types.BaseID         `json:"maintainersID"`
	Properties       []types.BaseProperty `json:"properties"`
	Lock             types.BaseHeight     `json:"lock"`
	Burn             types.BaseHeight     `json:"burn"`
}

var _ types.Request = (*request)(nil)

func RESTRequestHandler(cliContext context.CLIContext) http.HandlerFunc {

	makeBaseReqAndMsg := func(xprtRequest types.Request) (rest.BaseReq, sdkTypes.Msg) {
		request := xprtRequest.(request)

		baseReq := request.BaseReq
		from, Error := sdkTypes.AccAddressFromBech32(baseReq.From)
		if Error != nil {
			panic(errors.New(fmt.Sprintf("")))
		}

		noOfPropertiesSent := len(request.Properties)

		if noOfPropertiesSent > constants.MaxTraitCount {
			panic(errors.New(fmt.Sprintf("")))
		}

		var basePropertyList []types.BaseProperty
		for _, baseProperty := range request.Properties {
			basePropertyList = append(basePropertyList, baseProperty)
		}
		baseProperties := types.BaseProperties{BasePropertyList: basePropertyList}
		message := Message{
			From:             from,
			ChainID:          types.BaseID{IDString: baseReq.ChainID},
			MaintainersID:    request.MaintainersID,
			ClassificationID: request.ClassificationID,
			Properties:       &baseProperties,
			Lock:             request.Lock,
			Burn:             request.Burn,
		}
		return baseReq, message
	}

	var requestPrototype request

	return types.NewRESTRequest(requestPrototype).CreateRequest(cliContext, makeBaseReqAndMsg)
}
