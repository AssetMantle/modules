// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package order

import (
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

// type queryRequest struct {
//	ids.OrderID `json:"orderID" valid:"required~required field orderID missing"`
// }

var _ helpers.QueryRequest = (*QueryRequest)(nil)

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
func (queryRequest *QueryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if orderID, err := baseIDs.ReadOrderID(cliCommand.ReadString(constants.OrderID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(orderID), nil
	}
}
func (*QueryRequest) FromMap(vars map[string]string) (helpers.QueryRequest, error) {
	if orderID, err := baseIDs.ReadOrderID(vars[Query.GetName()]); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(orderID), nil
	}
}
func (queryRequest *QueryRequest) Encode() ([]byte, error) {
	return base.CodecPrototype().MarshalJSON(queryRequest)
}
func (queryRequest *QueryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	if err := base.CodecPrototype().UnmarshalJSON(bytes, queryRequest); err != nil {
		return nil, err
	}

	return queryRequest, nil
}
func requestPrototype() helpers.QueryRequest {
	return &QueryRequest{}
}
func queryRequestFromInterface(request helpers.QueryRequest) *QueryRequest {
	switch value := request.(type) {
	case *QueryRequest:
		return value
	default:
		return &QueryRequest{}
	}
}
func newQueryRequest(orderID ids.OrderID) helpers.QueryRequest {
	return &QueryRequest{OrderID: orderID.(*baseIDs.OrderID)}
}
