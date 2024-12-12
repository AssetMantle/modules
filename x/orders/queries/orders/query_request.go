// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package orders

import (
	"net/http"
	"strconv"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/orders/key"
)

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
	if err := queryRequest.Key.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf(err.Error())
	}

	if queryRequest.Limit > constants.PaginationLimit {
		return constants.InvalidRequest.Wrapf("limit cannot be greater than %d", constants.PaginationLimit)
	}

	if queryRequest.Limit < 0 {
		return constants.InvalidRequest.Wrapf("limit cannot be less than 0")
	}

	return nil
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	orderID, err := baseIDs.PrototypeOrderID().FromString(cliCommand.ReadString(constants.OrderID))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit := cliCommand.ReadInt(constants.Limit)

	return newQueryRequest(orderID.(ids.OrderID), int32(limit)), nil
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	orderID, err := baseIDs.PrototypeOrderID().FromString(httpRequest.URL.Query().Get(constants.Key.GetName()))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit, err := strconv.Atoi(httpRequest.URL.Query().Get(constants.Limit.GetName()))
	if err != nil {
		limit = query.DefaultLimit
	}

	return newQueryRequest(orderID.(ids.OrderID), int32(limit)), nil
}
func requestPrototype() helpers.QueryRequest {
	return &QueryRequest{}
}

func newQueryRequest(orderID ids.OrderID, limit int32) helpers.QueryRequest {
	return &QueryRequest{Key: key.NewKey(orderID).(*key.Key), Limit: limit}
}
