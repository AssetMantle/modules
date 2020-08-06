/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

//TODO make private
type Message struct {
	From          sdkTypes.AccAddress `json:"from" valid:"required~required field from missing matches(^cosmos[a-z0-9]{39}$)~invalid field from"`
	MaintainersID types.ID            `json:"maintainersID" valid:"required~required field maintainersID missing"`
	MakerID       types.ID            `json:"makerID" valid:"required~required field makerID missing"`
	TakerID       types.ID            `json:"takerID"`
	MakerSplit    sdkTypes.Dec        `json:"makerSplit" valid:"required~required field makerSplit missing"`
	MakerSplitID  types.ID            `json:"makerSplitID" valid:"required~required field makerSplitID missing"`
	ExchangeRate  sdkTypes.Dec        `json:"exchangeRate" valid:"required~required field exchangeRate missing"`
	TakerSplitID  types.ID            `json:"takerSplitID" valid:"required~required field takerSplitID missing"`
}

var _ sdkTypes.Msg = Message{}

func (message Message) Route() string { return Transaction.GetModuleName() }
func (message Message) Type() string  { return Transaction.GetName() }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(constants.IncorrectMessage, Error.Error())
	}
	return nil
}
func (message Message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(packageCodec.MustMarshalJSON(message))
}
func (message Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}

func messageFromInterface(msg sdkTypes.Msg) Message {
	switch value := msg.(type) {
	case Message:
		return value
	default:
		return Message{}
	}
}

func newMessage(from sdkTypes.AccAddress, maintainersID types.ID,
	fromID types.ID, toID types.ID, makerSplit sdkTypes.Dec, makerSplitID types.ID,
	exchangeRate sdkTypes.Dec, takerSplitID types.ID) sdkTypes.Msg {
	return Message{
		From:          from,
		MaintainersID: maintainersID,
		MakerID:       fromID,
		TakerID:       toID,
		MakerSplit:    makerSplit,
		MakerSplitID:  makerSplitID,
		ExchangeRate:  exchangeRate,
		TakerSplitID:  takerSplitID,
	}
}
