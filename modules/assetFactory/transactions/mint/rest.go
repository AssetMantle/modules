package mint

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type request struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	ClassificationID string       `json:"classificationID"`
	MaintainersID    string       `json:"maintainersID"`
	Properties       string       `json:"properties"`
	Lock             int          `json:"lock"`
	Burn             int          `json:"burn"`
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

		properties := strings.Split(request.Properties, constants.PropertiesSeparator)
		basePropertyList := make([]types.Property, 0)
		for _, property := range properties {
			traitIDAndProperty := strings.Split(property, constants.TraitIDAndPropertySeparator)
			if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
				basePropertyList = append(basePropertyList, types.NewProperty(types.NewID(traitIDAndProperty[0]), types.NewFact(traitIDAndProperty[1], types.NewSignatures(nil))))
			}
		}

		message := Message{
			From:             from,
			ChainID:          types.NewID(request.BaseReq.ChainID),
			MaintainersID:    types.NewID(request.MaintainersID),
			ClassificationID: types.NewID(request.ClassificationID),
			Properties:       types.NewProperties(basePropertyList),
			Lock:             types.NewHeight(request.Lock),
			Burn:             types.NewHeight(request.Burn),
		}
		return baseReq, message
	}

	var requestPrototype request

	return types.NewRESTRequest(requestPrototype).CreateRequest(cliContext, makeBaseReqAndMsg)
}
