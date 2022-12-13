// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package order

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/AssetMantle/modules/modules/orders/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type queryRequest struct {
	ids.OrderID `json:"orderID" valid:"required~required field orderID missing"`
}

var _ helpers.QueryRequest = (*queryRequest)(nil)

// Validate godoc
// @Summary Query order using order id
// @Description Able to query the order
// @Accept json
// @Produce json
// @Tags Orders
// @Param orderID path string true "order ID"
// @Success 200 {object} queryResponse "Message for a successful response"
// @Failure default  {object}  queryResponse "Message for an unexpected error response."
// @Router /orders/orders/{orderID} [get]
func (queryRequest queryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}
func (queryRequest) FromCLI(cliCommand helpers.CLICommand, _ context.CLIContext) (helpers.QueryRequest, error) {
	if orderID, err := baseIDs.ReadOrderID(cliCommand.ReadString(constants.OrderID)); err != nil {
		return queryRequest{}, err
	} else {
		return newQueryRequest(orderID), nil
	}
}
func (queryRequest) FromMap(vars map[string]string) (helpers.QueryRequest, error) {
	if orderID, err := baseIDs.ReadOrderID(vars[Query.GetName()]); err != nil {
		return queryRequest{}, err
	} else {
		return newQueryRequest(orderID), nil
	}
}
func (queryRequest queryRequest) Encode() ([]byte, error) {
	return common.Codec.MarshalJSON(queryRequest)
}
func (queryRequest queryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	if err := common.Codec.UnmarshalJSON(bytes, &queryRequest); err != nil {
		return nil, err
	}

	return queryRequest, nil
}
func requestPrototype() helpers.QueryRequest {
	return queryRequest{}
}
func queryRequestFromInterface(request helpers.QueryRequest) queryRequest {
	switch value := request.(type) {
	case queryRequest:
		return value
	default:
		return queryRequest{}
	}
}
func newQueryRequest(orderID ids.OrderID) helpers.QueryRequest {
	return queryRequest{OrderID: orderID}
}
